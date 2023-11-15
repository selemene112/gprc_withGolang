package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	netListen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	gRPCServer := grpc.NewServer()

	log.Printf("Server started at port :8080")
	if err := gRPCServer.Serve(netListen); err != nil {
		panic(err)
	}

}
