package main

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
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

	server := newGRPCServer() // FIXME
	pb.RegisterGetUsersServiceServer(server, &services.UserService{})
	reflection.Register(server)

	log.Infof("gRPC Server started: %s\n", serverHost)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}

func newGRPCServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	return s
}
