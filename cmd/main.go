package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zchelalo/sa_user/internal/server"
	"github.com/zchelalo/sa_user/pkg/config"
	"github.com/zchelalo/sa_user/pkg/sqlc/connection"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	conn, err := connection.NewConnection("postgres", source)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	ctx := context.Background()

	server.Run(ctx, config.Port, conn)
}
