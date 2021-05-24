package server

import (
	"fmt"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func ServeGRPC() {
	port := config.GetConfig().GRPCPort
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUsersServer(grpcServer, NewUsersGRPCServer())

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("User gRPC Service is listening on port %d", port)
	}()
}

type UsersGRPCServer struct {
	proto.UnimplementedUsersServer
}

func NewUsersGRPCServer() *UsersGRPCServer {
	grpcServer := &UsersGRPCServer{}
	return grpcServer
}

