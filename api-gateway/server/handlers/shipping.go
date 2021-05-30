package handlers

import (
	"net/http"
	"github.com/seb7887/go-microservices/clients"
)

const (
	ShippingAPIPath = "/api/v1/shipping"
)

func listShippingOrders() (map[string]interface{}, error) {
	shippingOrders, err := clients.ListShippingOrders()
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"shippingOrders": shippingOrders,
	}
	return resp, nil
}

func ListShippingOrders(w http.ResponseWriter, r *http.Request) {
	resp, err := listShippingOrders()
	if err != nil {
		HandleError(w, err.Error())
		return
	}
	PrepareResponse(w, resp)
}