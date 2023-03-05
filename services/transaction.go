package services

import (
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/adhityaf/bpjstk/models"
	"github.com/adhityaf/bpjstk/params"
	"github.com/adhityaf/bpjstk/repositories"
	"github.com/gin-gonic/gin"
)

type TransactionService struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: repo}
}

func (t *TransactionService) InsertDataTransaction(requests []params.InsertTransaction) *params.Response {
	start := time.Now()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	go t.dispatchWorkers(jobs, wg)
	readDataPerLineThenSendToWorker(requests, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	return &params.Response{
		Status: http.StatusCreated,
		Payload: gin.H{
			"message": "success",
			"Duration": fmt.Sprintf("done in %d seconds", int(math.Ceil(duration.Seconds()))),
		},
	}
}

func readDataPerLineThenSendToWorker(requests []params.InsertTransaction, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	rowOrdered := make([]interface{}, 0)
	for _, request := range requests {
		transaction := models.Transaction{
			Fullname: request.Fullname,
			Quantity: request.Quantity,
			Price:    request.Price,
		}

		rowOrdered = append(rowOrdered, transaction)
	}
	wg.Add(1)
	jobs <- rowOrdered
}

func (t *TransactionService) dispatchWorkers(jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		go func(jobs <-chan []interface{}, wg *sync.WaitGroup) {
			for job := range jobs {
				for _, job1 := range job{
					transaction := job1.(models.Transaction)
					t.transactionRepository.Create(&transaction)
				}
				wg.Done()
			}
		}(jobs, wg)
	}
}
