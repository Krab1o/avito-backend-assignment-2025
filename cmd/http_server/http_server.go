package main

import (
	"context"
	"log"

	apiAuth "github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth"
	apiBuying "github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying"
	apiTransaction "github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction"
	apiUser "github.com/Krab1o/avito-backend-assignment-2025/internal/api/user"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config/env"
	middlewareAuth "github.com/Krab1o/avito-backend-assignment-2025/internal/middleware/auth"
	middlewareParams "github.com/Krab1o/avito-backend-assignment-2025/internal/middleware/params"
	repositoryInventory "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory"
	repositoryMerch "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/merch"
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
	merchRepository := repositoryMerch.NewRepository(pool)

	// services
	transactionService := serviceTransaction.NewService(transactionRepository, userRepository)
	userService := serviceUser.NewService(userRepository)
	buyingService := serviceBuying.NewService(inventoryRepository, userRepository, merchRepository)
	authService := serviceAuth.NewHandler(userRepository, jwtConfig)

	// api
	transactionHandler := apiTransaction.NewHandler(transactionService)
	buyingHandler := apiBuying.NewHandler(buyingService)
	authHandler := apiAuth.NewHandler(authService)
	userHandler := apiUser.NewHandler(userService)

	serv := gin.Default()
	
	serv.POST(authPath, authHandler.Auth)

	securedEndpoints := serv.Group("")
	securedEndpoints.Use(middlewareAuth.JWTMiddleware(jwtConfig.Secret()))

	securedEndpoints.GET(infoPath, middlewareParams.NoParamsMiddleware(), userHandler.Info)
	securedEndpoints.POST(sendCoinPath, middlewareParams.NoParamsMiddleware(), transactionHandler.SendCoin)
	securedEndpoints.GET(buyItemPath, buyingHandler.Buy)

	//TODO: think of graceful shutdown handling
	if err := serv.Run(httpConfig.Address()); err != nil {
		log.Println("Server shutdown:", err)
	}
}