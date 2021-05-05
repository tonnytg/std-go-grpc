package services

import (
	"context"
	"fmt"
	"github.com/tonnytg/std-grpc-golang/pb/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	fmt.Println(request.Name)

	return &pb.User{
		Id: "123",
		Name: request.GetName(),
		Email: request.GetEmail(),
	}, nil
}