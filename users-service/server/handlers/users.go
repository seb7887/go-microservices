package handlers

import (
	"strconv"
	"log"
	"github.com/seb7887/go-microservices/db"
	"github.com/seb7887/go-microservices/models"
	"github.com/seb7887/go-microservices/helpers"
	"github.com/seb7887/go-microservices/proto"
	"github.com/seb7887/go-microservices/redis"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func Register(username string, email string, password string) *proto.UserResponse {
	hash := helpers.HashAndSalt([]byte(password))
	login := &models.Login{Hash: hash, Email: email}
	db.DB.Create(&login)

	user := &models.User{Username: username, Email: email}
	db.DB.Create(&user)

	response := &proto.UserResponse{
		UserId: strconv.FormatUint(uint64(user.ID), 10),
		Username: username,
		Email: email,
	}

	return response
}

func Login(email string, password string) (*proto.LoginResponse, error) {
	// Get user by email
	login := &models.Login{}
	db.DB.Where("email = ?", email).First(&login)

	// Verify Password
	isValid := helpers.VerifyPassword(password, login.Hash)

	// If password is valid, then generate JWT and return it in the response
	if isValid {
		user := &models.User{}
		db.DB.Where("email = ?", email).First(&user)
		userId := strconv.FormatUint(uint64(user.ID), 10)
		token := helpers.GenerateJWT(userId, email)
		err := redis.Set(token, []byte(userId))
		if err != nil {
			return nil, status.Error(codes.NotFound, "Redis error")
		}

		response := &proto.LoginResponse{Token: token}

		return response, nil
	} else {
		err := status.Error(codes.NotFound, "Wrong Password")
		return nil, err
	}
}

func GetUser(userId string) (*proto.UserResponse, error) {
	user := &models.User{}
	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		log.Fatal("Wrong convertion")
		er := status.Error(codes.NotFound, "Error")
		return nil, er
	}

	if db.DB.Where("id = ?", id).First(&user).RecordNotFound() {
		err := status.Error(codes.NotFound, "User not found")
		return nil, err
	}

	response := &proto.UserResponse{
		UserId: userId,
		Username: user.Username,
		Email: user.Email,
	}

	return response, nil
}