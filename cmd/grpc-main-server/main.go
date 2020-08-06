package main

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/mapserver2007/golang-example-app/grpc-web/common/constant"
	"github.com/mapserver2007/golang-example-app/grpc-web/common/database"
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

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterGetUsersServiceServer(server, &services.UserService{Connection: database.GetConnection()})
	pb.RegisterGetItemsServiceServer(server, &services.ItemService{Connection: database.GetConnection()})
	reflection.Register(server)

	log.Infof("gRPC Server(main) started: %s\n", serverHost)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
