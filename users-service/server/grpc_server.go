package server

import (
	"context"
	"fmt"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/proto"
	"github.com/seb7887/go-microservices/server/handlers"
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

func (gs *UsersGRPCServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	log.Println("CreateUser called")
	return handlers.Register(req.Username, req.Email, req.Password), nil
}

func (gs *UsersGRPCServer) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginResponse, error) {
	log.Println("LoginUser called")
	return handlers.Login(req.Email, req.Password)
}

func (gs *UsersGRPCServer) GetProfile(ctx context.Context, req *proto.GetProfileRequest) (*proto.UserResponse, error) {
	log.Println("GetProfile called")
	return handlers.GetUser(req.UserId)
}