package main

import (
	"log"
	"net/http"

	apiBuying "github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying"
	apiTransaction "github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config/env"
	serviceBuying "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying"
	serviceTransaction "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction"
	"github.com/gin-gonic/gin"
)

const (
	envPath = ".env"
	infoPath = "/api/info"
	sendCoinPath = "/api/sendCoin"
	buyItemPath = "/api/buy/:item"
)

const (
	errorParametersNotAllowed = "Query parameters are not allowed"
)

func rejectQueryParamsMiddleware(c *gin.Context) {
	if c.Request.URL.RawQuery != "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorParametersNotAllowed})
		c.Abort()
		return
	}
	c.Next()
}

func main() {
	err := config.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load config, %v", err)
	}
	httpConfig := env.NewHTTPConfig()
	pgConfig := env.NewPGConfig()
	transactionService := serviceTransaction.NewService()
	transactionHandler := apiTransaction.NewHandler(transactionService)
	buyingService := serviceBuying.NewService()
	buyingHandler := apiBuying.NewHandler(buyingService)
	log.Println(pgConfig.DSN())

	s := gin.Default()
	s.GET(infoPath, rejectQueryParamsMiddleware, transactionHandler.Info)
	s.POST(sendCoinPath, rejectQueryParamsMiddleware, transactionHandler.SendCoin)
	s.GET(buyItemPath, buyingHandler.Buy)
	s.Run(httpConfig.Address())
}