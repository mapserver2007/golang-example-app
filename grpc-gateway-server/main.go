package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	constant "github.com/mapserver2007/golang-example-app/common/constant"
	http_error "github.com/mapserver2007/golang-example-app/common/http"
	log "github.com/mapserver2007/golang-example-app/common/log"
	gw "github.com/mapserver2007/golang-example-app/gen/go"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	runtime.GlobalHTTPErrorHandler = http_error.CustomHTTPError

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("%s:%s", constant.GrpcMainServerHost, constant.GrpcMainServerPort)
	_ = gw.RegisterMainServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	gatewayEndpoint := fmt.Sprintf("%s:%s", constant.GatewayServerHost, constant.GatewayServerPort)
	log.Infof("gRPC Gateway Server started: %s\n", gatewayEndpoint)
	return http.ListenAndServe(gatewayEndpoint, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
