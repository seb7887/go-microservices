package main

import (
	"log"
	"github.com/seb7887/go-microservices/server"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/db"
)

func main() {
	port := config.GetConfig().Port
	
	// Initialize Database
	db.InitDatabase()
	db.AutoMigrate()

	log.Printf("Starting server in port %d", port)
	server.ServeGRPC()
	err := server.Serve(port)
	if err != nil {
		log.Fatal(err)
	}
}