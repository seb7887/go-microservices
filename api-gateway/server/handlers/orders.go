package handlers

import (
	"net/http"
	"github.com/seb7887/go-microservices/clients"
)

const (
	OrdersAPIPath = "/api/v1/orders"
)

type CreateOrderRequest struct {
	ProductName string
	TotalAmount int32
}

func createOrder(userId string, productName string, totalAmount int32) (map[string]interface{}, error) {
	order, err := clients.CreateOrder(userId, productName, totalAmount)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"orderId": order.OrderId,
		"userId": userId,
		"productName": productName,
		"totalAmount": totalAmount,
	}
	return resp, nil
}

func listOrders() (map[string]interface{}, error) {
	orders, err := clients.ListOrders()
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"orders": orders,
	}
	return resp, nil
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")
	body, err := ReadBody(r)
	if err != nil {
		HandleError(w, "Empty body")
		return
	}
	var parsedBody CreateOrderRequest
	err = ParseBody(body, &parsedBody)
	if err != nil {
		HandleError(w, "Error parsing body")
		return
	}

	resp, err := createOrder(userId, parsedBody.ProductName, parsedBody.TotalAmount)
	if err != nil {
		HandleError(w, err.Error())
		return
	}
	PrepareResponse(w, resp)
}

func ListOrders(w http.ResponseWriter, r *http.Request) {
	resp, err := listOrders()
	if err != nil {
		HandleError(w, err.Error())
		return
	}
	PrepareResponse(w, resp)
}