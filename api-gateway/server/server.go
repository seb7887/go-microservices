package server

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/seb7887/go-microservices/server/handlers"
)

var getUserProfileHandler = http.HandlerFunc(handlers.GetUserProfile)

func Serve(port int) error {
	router := mux.NewRouter()
	router.Use(handlers.PanicHandler)
	router.HandleFunc(handlers.HealtAPIPath, handlers.Health).Methods("GET")
	router.HandleFunc(handlers.UsersAPIPath, handlers.SignUp).Methods("POST")
	router.HandleFunc(handlers.LoginAPIPath, handlers.SignIn).Methods("POST")
	router.HandleFunc(handlers.UserAPIPath, handlers.AuthMiddleware(getUserProfileHandler)).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}