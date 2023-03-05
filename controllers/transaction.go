package controllers

import (
	"net/http"

	"github.com/adhityaf/bpjstk/params"
	"github.com/adhityaf/bpjstk/services"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: *transactionService}
}

func (t *TransactionController) InsertDataTransaction(ctx *gin.Context) {
	var req []params.InsertTransaction

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}
	result := t.transactionService.InsertDataTransaction(req)

	ctx.JSON(result.Status, result.Payload)
}
