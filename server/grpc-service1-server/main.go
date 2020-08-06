package main

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/mapserver2007/golang-example-app/server/common/constant"
	"github.com/mapserver2007/golang-example-app/server/common/database"
	"github.com/mapserver2007/golang-example-app/server/common/log"
	pb "github.com/mapserver2007/golang-example-app/server/gen/go"
	"github.com/mapserver2007/golang-example-app/server/grpc-service1-server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
