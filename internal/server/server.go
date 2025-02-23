package server

import (
	"context"
	"fmt"
	"net"

	_ "github.com/lib/pq"
	userInfrastructure "github.com/zchelalo/sa_user/internal/modules/user/infrastructure"
	"github.com/zchelalo/sa_user/pkg/bootstrap"
	"github.com/zchelalo/sa_user/pkg/proto"
	userDb "github.com/zchelalo/sa_user/pkg/sqlc/user/db"
	"github.com/zchelalo/sa_user/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() {
	logger := bootstrap.GetLogger()

	conn, err := bootstrap.GetInstance()
	if err != nil {
		logger.Fatal("cannot connect to db:", err)
	}

	ctx := context.Background()

	userStore := userDb.NewStore(conn)
	userRouter := userInfrastructure.NewUserRouter(userStore, ctx)

	config := util.GetConfig()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.Fatal("cannot listen:", err)
	}

	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, userRouter)

	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		logger.Fatalf("Error serving: %s", err.Error())
	}

	logger.Println("Server running on port:", config.Port)
}
