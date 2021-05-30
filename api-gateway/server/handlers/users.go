package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/seb7887/go-microservices/helpers"
	"github.com/seb7887/go-microservices/clients"
)

type SignUpRequest struct {
	Username string
	Email string
	Password string
}

type SignInRequest struct {
	Email string
	Password string
}

type GetUserProfileRequest struct {
	UserId string
}

const (
	UsersAPIPath = "/api/v1/users"
	LoginAPIPath = "/api/v1/login"
	UserAPIPath = "/api/v1/users/{id}"
)

func createUser(username string, email string, password string) (map[string]interface{}, error) {
	valid := helpers.Validate(
		[]helpers.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
		},
	)

	if valid {
		user, err := clients.CreateUser(username, email, password)
		if err != nil {
			return nil, err
		}
		resp := map[string]interface{}{
			"userId": user.UserId,
			"username": user.Username,
			"email": user.Email,
		}
		return resp, nil
	} else {
		return nil, fmt.Errorf("Invalid body format")
	}
}

func loginUser(email string, password string) (map[string]interface{}, error) {
	valid := helpers.Validate(
		[]helpers.Validation{
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
		},
	)

	if valid {
		token, err := clients.LoginUser(email, password)
		if err != nil {
			return nil, err
		}
		resp := map[string]interface{}{
			"token": token,
		}
		return resp, nil
	} else {
		return nil, fmt.Errorf("Invalid body format")
	}
}

func getUserProfile(userId string) (map[string]interface{}, error) {
	user, err := clients.GetProfile(userId)
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{
		"userId": user.UserId,
		"username": user.Username,
		"email": user.Email,
	}
	return resp, nil
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ReadBody(r)
	if err != nil {
		HandleError(w, "Empty body")
		return
	}
	var parsedBody SignUpRequest
	err = ParseBody(body, &parsedBody)
	if err != nil {
		HandleError(w, "Error parsing body")
		return
	}

	resp, err := createUser(parsedBody.Username, parsedBody.Email, parsedBody.Password)
	if err != nil {
		HandleError(w, err.Error())
		return
	}
	PrepareResponse(w, resp)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := ReadBody(r)
	if err != nil {
		HandleError(w, "Empty Body")
		return
	}
	var parsedBody SignInRequest
	err = ParseBody(body, &parsedBody)
	if err != nil {
		HandleError(w, "Error parsing body")
		return
	}

	resp, err := loginUser(parsedBody.Email, parsedBody.Password)
	if err != nil {
		HandleError(w, err.Error())
		return
	}
	PrepareResponse(w, resp)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := r.Header.Get("userId")

	if id != userId {
		HandleError(w, "Unauthorized")
		return
	}
	
	resp, err := getUserProfile(userId)
	if err != nil {
		HandleError(w, err.Error())
		return
	}
	PrepareResponse(w, resp)
}