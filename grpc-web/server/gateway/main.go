package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	config "github.com/mapserver2007/golang-example-app/grpc-web/config"
	gw "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("%s:%s", config.Host, config.GrpcServerPort)
	err := gw.RegisterGetUsersServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	gatewayEndpoint := fmt.Sprintf("%s:%s", config.Host, config.GatewayServerPort)
	log.Printf("gRPC Gateway Server started: %s\n", gatewayEndpoint)
	return http.ListenAndServe(gatewayEndpoint, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
