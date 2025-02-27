package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	userApplication "github.com/zchelalo/sa_user/internal/modules/user/application"
	userGRPC "github.com/zchelalo/sa_user/internal/modules/user/infrastructure/adapters/grpc"
	userPostgresRepo "github.com/zchelalo/sa_user/internal/modules/user/infrastructure/repositories/postgres"
	"github.com/zchelalo/sa_user/internal/server"
	"github.com/zchelalo/sa_user/pkg/bootstrap"
	"github.com/zchelalo/sa_user/pkg/proto"
	"github.com/zchelalo/sa_user/pkg/sqlc/db"
	"github.com/zchelalo/sa_user/pkg/util"
	"google.golang.org/grpc"
)

func main() {
	logger := bootstrap.GetLogger()

	config, err := util.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	conn, err := bootstrap.GetInstance("postgres", source)
	if err != nil {
		logger.Fatal("cannot connect to db:", err)
	}

	dbStore := db.New(conn)

	userRepository := userPostgresRepo.New(dbStore)
	userUseCases := userApplication.New(userRepository)
	userRouter := userGRPC.New(userUseCases)

	server := server.New(config.Port,
		func(s *grpc.Server) { proto.RegisterUserServiceServer(s, userRouter) },
	)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		logger.Println("shutting down gracefully...")
		bootstrap.Close()
		os.Exit(0)
	}()

	server.Start()
}
