package server

import (
	"fmt"
	"net"

	"github.com/zchelalo/sa_user/pkg/bootstrap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	port       int32
}

func New(port int32, serviceRegistrations ...func(*grpc.Server)) *Server {
	server := grpc.NewServer()

	for _, register := range serviceRegistrations {
		register(server)
	}

	reflection.Register(server)

	return &Server{
		grpcServer: server,
		port:       port,
	}
}

func (s *Server) Start() {
	logger := bootstrap.GetLogger()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		logger.Fatal("cannot listen:", err)
	}

	if err := s.grpcServer.Serve(listen); err != nil {
		logger.Fatalf("Error serving: %s", err.Error())
	}

	logger.Println("Server running on port:", s.port)
}
