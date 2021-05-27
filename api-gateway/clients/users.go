package clients

import (
	"fmt"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/seb7887/go-microservices/config"
	"github.com/seb7887/go-microservices/proto"
)

type User struct {
	UserId string
	Username string
	Email string
}

func CreateUser(username string, email string, password string) (*User, error) {
	addr := fmt.Sprintf("%s:%s", config.GetConfig().UsersHost, config.GetConfig().UsersPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.CreateUserRequest{
		Username: username,
		Email: email,
		Password: password,
	}
	res, err := client.CreateUser(ctx, &req)
	if err != nil {
		return nil, err
	}

	user := User{
		UserId: res.GetUserId(),
		Username: res.GetUsername(),
		Email: res.GetEmail(),
	}

	return &user, nil
}

func LoginUser(email string, password string) (*string, error) {
	addr := fmt.Sprintf("%s:%s", config.GetConfig().UsersHost, config.GetConfig().UsersPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.LoginUserRequest{
		Email: email,
		Password: password,
	}
	res, err := client.LoginUser(ctx, &req)
	if err != nil {
		return nil, err
	}
	token := res.GetToken()

	return &token, nil
}

func GetProfile(userId string) (*User, error) {
	addr := fmt.Sprintf("%s:%s", config.GetConfig().UsersHost, config.GetConfig().UsersPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.GetProfileRequest{
		UserId: userId,
	}
	res, err := client.GetProfile(ctx, &req)
	if err != nil {
		return nil, err
	}
	user := User{
		UserId: res.GetUserId(),
		Username: res.GetUsername(),
		Email: res.GetEmail(),
	}

	return &user, nil
}