package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	userInfrastructure "github.com/zchelalo/sa_user/internal/modules/user/infrastructure"
	userProto "github.com/zchelalo/sa_user/pkg/proto/user"
	userDb "github.com/zchelalo/sa_user/pkg/sqlc/user/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, port int32, conn *sql.DB) {
	userStore := userDb.NewStore(conn)
	userRouter := userInfrastructure.NewUserRouter(userStore, ctx)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	server := grpc.NewServer()
	userProto.RegisterUserServiceServer(server, userRouter)

	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}

	fmt.Println("Server running on port:", port)
}
