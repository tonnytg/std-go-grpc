package main

import (
	"context"
	"fmt"
	"github.com/tonnytg/std-grpc-golang/pb/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main()  {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	AddUser(client)
	AddUserVerbose(client)
}

func AddUser( client pb.UserServiceClient)  {
	req := &pb.User{
		Id: "0",
		Name: "Antonio",
		Email: "tonnytg@gmail.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatal("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient)  {

	req := &pb.User{
		Id: "0",
		Name: "Joao",
		Email: "tonnytg@gmail.com",
	}

	responseSteam, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf( "Could not make gRPC request %v", err)
	}

	for {
		stream, err := responseSteam.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())

	}
	
}
