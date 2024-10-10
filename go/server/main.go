package main

import (
	"context"
	"fmt"
	"grpccode/configs"
	"grpccode/database"
	"grpccode/database/entities"
	"grpccode/request"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserFeaturesServer struct {
	request.UnimplementedUserFeaturesServer
}

func (s *UserFeaturesServer) RegisterUser(ctx context.Context, req *request.Person) (*request.Response, error) {
	t, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	log.Println(req)

	reqEntity := &entities.Users{
		Name:        req.Name,
		Email:       req.Email,
		DateOfBirth: t,
		Username:    req.Username,
		Password:    req.Password,
	}

	if err := database.DB.Where("username = ? OR email = ?", reqEntity.Username, reqEntity.Email).First(&reqEntity).Error; err == nil {
		message := "Cannot create this user as it already exists"

		return &request.Response{
			Message: message,
			Person:  nil,
		}, fmt.Errorf("user already exists")
	}

	if err := database.DB.Create(&reqEntity).Error; err != nil {
		message := "Cannot create this user, unexpected error"

		return &request.Response{
			Message: message,
			Person:  nil,
		}, fmt.Errorf("cannot create this user")
	}
	message := "Create user success"
	return &request.Response{
		Message: message,
		Person:  req,
	}, nil
}

func (s *UserFeaturesServer) Login(ctx context.Context, req *request.LoginRequest) (*request.Response, error) {
	var user *entities.Users
	if err := database.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		log.Println(err)
		return &request.Response{
			Message: "Username or Password is wrong",
			Person:  nil,
		}, err
	}

	return &request.Response{
		Message: "Welcome, " + user.Name,
		Person: &request.Person{
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func main() {
	configs.InitializeEnv()
	database.ConnectDB()
	serv := grpc.NewServer()
	var userFeaturesServer UserFeaturesServer
	request.RegisterUserFeaturesServer(serv, &userFeaturesServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(serv)
	log.Println("starting server...")
	log.Fatal(serv.Serve(lis))
}
