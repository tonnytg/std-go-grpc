package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Could not connect: %v", err)
	}

	grpcServer = new grpc.NewServer()

	if err := gprcServer.Serve(listen; err != nil {
		log.Fatal("Could not connect %v:", err)
	}

}