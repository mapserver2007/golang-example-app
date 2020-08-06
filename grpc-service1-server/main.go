package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/mapserver2007/golang-example-app/common/constant"
	"github.com/mapserver2007/golang-example-app/common/database"
	"github.com/mapserver2007/golang-example-app/common/log"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service1-server/services"
)

func main() {
	serverHost := fmt.Sprintf("%s:%s", constant.GrpcService1ServerHost, constant.GrpcService1ServerPort)
	listen, err := net.Listen("tcp", serverHost)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterUserServiceServer(server, &services.UserService{Connection: database.GetConnection()})
	reflection.Register(server)

	log.Infof("gRPC Service1 Server started: %s\n", serverHost)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
