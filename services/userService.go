package services

import (
	"context"
	"fmt"
	"github.com/tonnytg/std-grpc-golang/pb/pb"
	"time"
)

//type UserServiceClient interface {
//	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
//	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
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

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error  {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:  &pb.User{},
	})

	time.Sleep( time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:  &pb.User{},
	})

	time.Sleep( time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User:  &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep( time.Second * 3)

	return nil
}