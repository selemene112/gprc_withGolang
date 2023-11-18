package main

import (
	"gprc_protocolbuffer/cmd/Config"
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

	db := Config.StartPostgresqlDB()

	gRPCServer := grpc.NewServer()
	UserService := services.UserServices{DB: db}
	UsersPb.RegisterProductServiceServer(gRPCServer, &UserService)

	log.Printf("Server started at port :8080")
	if err := gRPCServer.Serve(netListen); err != nil {
		panic(err)
	}

}
