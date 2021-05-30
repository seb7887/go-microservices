package server

import (
	"fmt"
	"net/http"
	"github.com/seb7887/go-microservices/server/handlers"
	"github.com/seb7887/go-microservices/rabbitmq"
)

func Serve(port int) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/health", handlers.Health)

	rabbitmq.AMQPListen()
	return http.ListenAndServe(fmt.Sprintf(":%d", port), serveMux)
}