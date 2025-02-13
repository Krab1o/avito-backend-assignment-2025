package main

import (
	"context"
	"log"
	"net/http"

	apiAuth "github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth"
	apiBuying "github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying"
	apiTransaction "github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction"
	apiUser "github.com/Krab1o/avito-backend-assignment-2025/internal/api/user"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config/env"
	repositoryInventory "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory"
	repositoryTransaction "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction"
	repositoryUser "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user"
	serviceAuth "github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth"
	serviceBuying "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying"
	serviceUser "github.com/Krab1o/avito-backend-assignment-2025/internal/service/info"
	serviceTransaction "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	envPath = ".env"
	infoPath = "/api/info"
	sendCoinPath = "/api/sendCoin"
	buyItemPath = "/api/buy/:item"
	authPath = "/api/auth"
)

const (
	errorParametersNotAllowed = "Query parameters are not allowed"
)

func checkParamsMiddleware(c *gin.Context) {
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
	httpConfig, err := env.NewHTTPConfig()
	if err != nil {
		log.Fatalf("Failed to load http config, %v", err)
	}
	jwtConfig, err := env.NewJWTConfig()
	if err != nil {
		log.Fatalf("Failed to load jwt config, %v", err)
	}
	pgConfig, err := env.NewPGConfig()
	log.Println(jwtConfig)
	//TODO: REFACTOR
	// POSTGRES CODE STARTED
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("DB is not reachable, %v", err)
	}

	// repositories
	transactionRepository := repositoryTransaction.NewRepository(pool)
	userRepository := repositoryUser.NewRepository(pool)
	inventoryRepository := repositoryInventory.NewRepository(pool)

	// services
	transactionService := serviceTransaction.NewService(transactionRepository)
	userService := serviceUser.NewService(userRepository)
	buyingService := serviceBuying.NewService(inventoryRepository)
	authService := serviceAuth.NewHandler(userRepository, jwtConfig)

	// api
	transactionHandler := apiTransaction.NewHandler(transactionService)
	buyingHandler := apiBuying.NewHandler(buyingService)
	authHandler := apiAuth.NewHandler(authService)
	userHandler := apiUser.NewHandler(userService)

	s := gin.Default()
	
	//DO NOT ADD MIDDLEWARE
	s.POST(authPath, authHandler.Auth)

	//ADD AUTH MIDDLEWARE TO THIS ENDPOINTS
	s.GET(infoPath, checkParamsMiddleware, userHandler.Info)
	s.POST(sendCoinPath, checkParamsMiddleware, transactionHandler.SendCoin)
	s.GET(buyItemPath, buyingHandler.Buy)

	//TODO: think of graceful shutdown handling
	if err := s.Run(httpConfig.Address()); err != nil {
		log.Println("Server shutdown:", err)
	}
}