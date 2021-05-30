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
	proto.RegisterOrdersServer(grpcServer, NewOrdersGRPCServer())

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("Order gRPC Service is listening on port %d", port)
	}()
}

type OrdersGRPCServer struct {
	proto.UnimplementedOrdersServer
}

func NewOrdersGRPCServer() *OrdersGRPCServer {
	grpcServer := &OrdersGRPCServer{}
	return grpcServer
}

func (gs *OrdersGRPCServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	log.Println("CreateOrder called")
	return handlers.Create(req.UserId, req.ProductName, req.TotalAmount)
}

func (gs *OrdersGRPCServer) GetOrder(ctx context.Context, req *proto.GetOrdersRequest) (*proto.OrdersResponse, error) {
	log.Println("GetOrders called")
	return handlers.ListOrders(req.Limit, req.Offset)
}