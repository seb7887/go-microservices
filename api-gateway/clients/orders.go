package clients

import (
	"fmt"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/proto"
)

type Order struct {
	OrderId string
}

func CreateOrder(userId string, productName string, totalAmount int32) (*Order, error) {
	addr := fmt.Sprintf("%s:%s", config.GetConfig().OrdersHost, config.GetConfig().OrdersPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewOrdersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.CreateOrderRequest{
		UserId: userId,
		ProductName: productName,
		TotalAmount: totalAmount,
	}
	res, err := client.CreateOrder(ctx, &req)
	if err != nil {
		return nil, err
	}

	order := Order{OrderId: res.GetOrderId()}

	return &order, nil
}

func ListOrders() ([]*proto.OrderResponse, error) {
	addr := fmt.Sprintf("%s:%s", config.GetConfig().OrdersHost, config.GetConfig().OrdersPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewOrdersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.GetOrdersRequest{
		Limit: 10,
		Offset: 10,
	}
	res, err := client.GetOrder(ctx, &req)
	if err != nil {
		return nil, err
	}

	orders := res.GetOrders()
	return orders, nil
}

