package main

import (
	"log"

	"github.com/adhityaf/bpjstk/config"
	"github.com/adhityaf/bpjstk/controllers"
	"github.com/adhityaf/bpjstk/repositories"
	"github.com/adhityaf/bpjstk/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.GetPostgresDB()
	if err != nil {
		log.Fatal("Error connect DB")
	}
	route := gin.Default()

	transactionRepository := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	transactionController := controllers.NewTransactionController(transactionService)

	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/insertdata", transactionController.InsertDataTransaction)
	}

	route.Run()
}
