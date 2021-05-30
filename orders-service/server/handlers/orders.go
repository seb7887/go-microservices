package handlers

import (
	"strconv"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/seb7887/go-microservices/db"
	"github.com/seb7887/go-microservices/models"
	"github.com/seb7887/go-microservices/proto"
	"github.com/seb7887/go-microservices/rabbitmq"
)

func Create(userId string, productName string, totalAmount int32) (*proto.OrderResponse, error) {
	order := &models.Order{UserId: userId, ProductName: productName, TotalAmount: totalAmount}
	result := db.DB.Create(&order)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	// Publish a shipping request to the message broker
	err := rabbitmq.PublishMessage(order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &proto.OrderResponse{OrderId: strconv.FormatUint(uint64(order.ID), 10)}
	return response, nil
}

func ListOrders(limit int64, offset int64) (*proto.OrdersResponse, error) {
	dbOrders := &[]models.Order{}
	result := db.DB.Limit(limit).Find(&dbOrders)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var orders []*proto.OrderResponse
	for _, order := range *dbOrders {
		orders = append(orders, &proto.OrderResponse{OrderId: strconv.FormatUint(uint64(order.ID), 10)})
	}

	response := &proto.OrdersResponse{
		Orders: orders,
	}

	return response, nil
}