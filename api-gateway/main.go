package main

import (
	"log"
	"github.com/seb7887/go-microservices/server"
	"github.com/seb7887/go-microservices/config"
)

func main() {
	port := config.GetConfig().Port	
	log.Printf("Starting server in port %d", port)
	err := server.Serve(port)
	if err != nil {
		log.Fatal(err)
	}
}