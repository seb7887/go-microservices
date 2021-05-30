package clients

import (
	"fmt"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/proto"
)

func ListShippingOrders() ([]*proto.ShippingOrderResponse, error) {
	addr := fmt.Sprintf("%s:%s", config.GetConfig().ShippingHost, config.GetConfig().ShippingPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewShippingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.GetShippingRequest{
		Limit: 10,
		Offset: 10,
	}
	res, err := client.GetShippingOrders(ctx, &req)
	if err != nil {
		return nil, err
	}

	orders := res.GetShippingOrders()
	return orders, nil
}

