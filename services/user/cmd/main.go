package main

import (
	"fmt"
	gen "gen/user"
	"learning_golang/user/internal/config"
	grpc2 "learning_golang/user/internal/handler/grpc"
	"learning_golang/user/internal/repository/postgres"
	"learning_golang/user/internal/service"
	"learning_golang/user/pkg/db"
	"log"
	"logger"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config")
	}

	logger, err := logger.NewLogger("UserService")

	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	logger.Info("Running microservice on http://localhost:%s/", cfg.ServerPort)

	postgresCfg := &db.PostgresConfig{}

	postgresDB, err := db.NewPostgresDB(postgresCfg, logger)

	if err != nil {
		logger.Error("Error initializing postgres db: %v", err)
	}

	repo := postgres.NewUserRepo(postgresDB, logger)
	service := service.NewUserService(repo, logger)

	logger.Info("Complete initializing user repository")

	grpcServer := grpc.NewServer()
	userServer := grpc2.NewUserServer(service, logger)

	gen.RegisterUserServiceServer(grpcServer, userServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.ServerPort))

	if err != nil {
		logger.Fatal("Error listening on port %v: %v", cfg.ServerPort, err)
	}

	go func() {

		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Error serving grpc server on port %v: %v", cfg.ServerPort, err)
		}

	}()

	logger.Info("Server listening on port %v", cfg.ServerPort)
}
