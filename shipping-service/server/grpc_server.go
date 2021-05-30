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
	proto.RegisterShippingServer(grpcServer, NewShippingGRPCServer())

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("Shipping gRPC Service is listening on port %d", port)
	}()
}

type ShippingGRPCServer struct {
	proto.UnimplementedShippingServer
}

func NewShippingGRPCServer() *ShippingGRPCServer {
	grpcServer := &ShippingGRPCServer{}
	return grpcServer
}

func (gs *ShippingGRPCServer) GetShippingOrders(ctx context.Context, req *proto.GetShippingRequest) (*proto.ShippingOrders, error) {
	log.Println("GetShippingOrders called")
	return handlers.ListShippingOrders(req.Limit, req.Offset)
}