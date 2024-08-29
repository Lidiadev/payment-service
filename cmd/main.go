package main

//import (
//	"context"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	swaggerfiles "github.com/swaggo/files"
//	ginSwagger "github.com/swaggo/gin-swagger"
//	"log"
//	"payment-service/internal/application/usecases"
//	"payment-service/internal/domain"
//	"payment-service/internal/infrastructure/config"
//	"payment-service/internal/infrastructure/gateways"
//	"payment-service/internal/infrastructure/logging"
//	"payment-service/internal/infrastructure/persistence"
//	"payment-service/internal/infrastructure/resilience"
//	"payment-service/internal/infrastructure/workers"
//	"payment-service/internal/interfaces/handlers"
//	"payment-service/internal/interfaces/validators"
//	"payment-service/pkg"
//	"time"
//)
//
//func main() {
//	r := gin.Default()
//
//	env := "development"
//
//	logger, err := logging.New(env)
//	if err != nil {
//		fmt.Printf("Error initializing logging: %v", err)
//		return
//	}
//
//	cfg, err := config.InitConfig()
//	if err != nil {
//		fmt.Printf("Error initializing configuration: %v", err)
//		return
//	}
//
//	gatewayAConfig := cfg.Gateways["gatewayA"]
//	gatewayBConfig := cfg.Gateways["gatewayB"]
//
//	resilience.ConfigureHystrix()
//	resilience.InitializeHystrixCommand(resilience.HystrixConfig{
//		Name:                   "gatewayA-deposit",
//		Timeout:                gatewayAConfig.Timeout,
//		ErrorPercentThreshold:  50,
//		RequestVolumeThreshold: 5,
//		SleepWindow:            10000,
//	})
//
//	gatewayAClient := gateways.NewGatewayAClient(gatewayAConfig.URL, "gatewayA-deposit")
//	gatewayBClient := gateways.NewGatewayBClient(gatewayBConfig.URL, time.Duration(gatewayBConfig.Timeout)*time.Millisecond, gatewayBConfig.MaxRetries)
//
//	gatewayA := gateways.NewGatewayA(gatewayAClient)
//	gatewayB := gateways.NewGatewayB(gatewayBClient)
//
//	gatewayFactory := gateways.NewGatewayFactory()
//	gatewayFactory.RegisterGateway(domain.GatewayAName, gatewayA)
//	gatewayFactory.RegisterGateway(domain.GatewayBName, gatewayB)
//
//	eventStore := persistence.NewInMemoryEventStore(logger)
//	queue := pkg.NewQueue()
//
//	transactionService := usecases.NewTransactionService(eventStore, queue)
//
//	validator := validators.NewRequestValidator()
//	depositHandler := handlers.NewDepositHandler(transactionService, validator, logger)
//
//	worker := workers.NewEventWorker(queue, eventStore, gatewayFactory)
//
//	// Start worker in a separate goroutine
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	go func() {
//		if err := worker.Start(ctx); err != nil {
//			log.Fatalf("Worker failed: %v", err)
//		}
//	}()
//
//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
//
//	r.POST("/v1/deposit", depositHandler.HandleDeposit)
//
//	logger.Info("Starting payment service on port 8080...")
//	if err := r.Run(":8080"); err != nil {
//		logger.Fatal("could not start server: %v\n" + err.Error())
//	}
//}
