package main

import (
	"log"
	"net"

	services "gprc_protocolbuffer/cmd/Services"
	UsersPb "gprc_protocolbuffer/pkg"

	"google.golang.org/grpc"
)

func main() {

	netListen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	gRPCServer := grpc.NewServer()
	UserService := services.UserServices{}
	UsersPb.RegisterProductServiceServer(gRPCServer, &UserService)

	log.Printf("Server started at port :8080")
	if err := gRPCServer.Serve(netListen); err != nil {
		panic(err)
	}

}
