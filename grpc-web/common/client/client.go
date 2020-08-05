package client

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	"google.golang.org/grpc"
)

func GrpcClient(address string, ctx context.Context, in *empty.Empty) (interface{}, error) {
	// TODO insecureを直す、inはenptyだとだめなのでinterfaceとかで受ける
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewGetItemsServiceClient(conn)
	return c.GetItems(ctx, in)
}
