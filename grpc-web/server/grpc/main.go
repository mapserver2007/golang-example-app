package main

import (
	"fmt"
	"log"
	"net"

	config "github.com/mapserver2007/golang-example-app/grpc-web/config"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	services "github.com/mapserver2007/golang-example-app/grpc-web/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	serverHost := fmt.Sprintf("%s:%s", config.Host, config.GrpcServerPort)
	listen, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	pb.RegisterGetUsersServiceServer(server, &services.UserService{})
	reflection.Register(server)

	log.Printf("gRPC Server started: %s\n", serverHost)

	if err := server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
