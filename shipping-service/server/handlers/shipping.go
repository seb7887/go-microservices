package handlers

import (
	"strconv"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/seb7887/go-microservices/db"
	"github.com/seb7887/go-microservices/models"
	"github.com/seb7887/go-microservices/proto"
)

func ListShippingOrders(limit int64, offset int64) (*proto.ShippingOrders, error) {
	dbShippings := &[]models.Shipping{}
	result := db.DB.Limit(limit).Find(&dbShippings)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var shippings []*proto.ShippingOrderResponse
	for _, shipping := range *dbShippings {
		shippings = append(shippings, &proto.ShippingOrderResponse{
			ShippingId: strconv.FormatUint(uint64(shipping.ID), 10),
			UserId: shipping.UserId,
			OrderId: shipping.OrderId,
		})
	}

	response := &proto.ShippingOrders{
		ShippingOrders: shippings,
	}

	return response, nil
}