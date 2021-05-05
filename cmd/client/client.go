package main

import (
	"context"
	"fmt"
	"github.com/tonnytg/std-grpc-golang/pb/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main()  {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUsers(client pb.UserServiceClient)  {
	reqs := []*pb.User{
		&pb.User{
			Id: "1",
			Name: "tonnyt1",
			Email: "tonnytg@gmail.com",
		},
		&pb.User{
			Id: "2",
			Name: "tonnyt2",
			Email: "tonnytg@gmail.com",
		},
		&pb.User{
			Id: "3",
			Name: "tonnyt3",
			Email: "tonnytg@gmail.com",
		},
		&pb.User{
			Id: "4",
			Name: "tonnyt4",
			Email: "tonnytg@gmail.com",
		},
		&pb.User{
			Id: "5",
			Name: "tonnyt5",
			Email: "tonnytg@gmail.com",
		},
		&pb.User{
			Id: "6",
			Name: "tonnyt6",
			Email: "tonnytg@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
	
}

func AddUserStreamBoth(client pb.UserServiceClient)  {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id: "1",
			Name: "tonnyt1",
			Email: "tonnytg@gmail.com",
		},
		&pb.User{
			Id: "2",
			Name: "tonnyt2",
			Email: "tonnytg@gmail.com",
		},
	}

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	wait := make(chan int)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}

			fmt.Printf("Recebendo user %v com status %v \n", res.GetUser().GetName(), res.GetStatus() )
		}

		close(wait)
	}()

	<- wait

}
