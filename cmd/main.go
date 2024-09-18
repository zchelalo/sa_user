package main

import (
	"fmt"

	"github.com/zchelalo/sa_user/internal/server"
	"github.com/zchelalo/sa_user/pkg/bootstrap"
	"github.com/zchelalo/sa_user/pkg/util"
)

func main() {
	bootstrap.InitLogger()
	logger := bootstrap.GetLogger()

	config, err := util.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	err = bootstrap.InitInstance("postgres", source)
	if err != nil {
		logger.Fatal("cannot connect to db:", err)
	}

	server.Start()
}
