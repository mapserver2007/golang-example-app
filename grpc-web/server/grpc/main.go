package main

import (
	"fmt"
	"net"

	"github.com/mapserver2007/golang-example-app/grpc-web/common/constant"
	"github.com/mapserver2007/golang-example-app/grpc-web/common/log"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-web/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	serverHost := fmt.Sprintf("%s:%s", constant.GrpcServerHost, constant.GrpcServerPort)
	listen, err := net.Listen("tcp", serverHost)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterGetUsersServiceServer(server, &services.UserService{})
	reflection.Register(server)

	log.Infof("gRPC Server started: %s\n", serverHost)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
