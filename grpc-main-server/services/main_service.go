package services

import (
	"context"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mapserver2007/golang-example-app/common/log"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
)

type MainService struct {
}

func (s *MainService) GetUsersAndItems(ctx context.Context, in *empty.Empty) (*pb.GetUsersAndItemsResponse, error) {
	tx := newSagaService(ctx, "grpc-main-server")
	tx.start()
	users := s.grpcService1Clinet(ctx, in).Users
	items := s.grpcService2Clinet(ctx, in).Items
	tx.end()

	return &pb.GetUsersAndItemsResponse{Users: users, Items: items}, nil
}

func (s *MainService) PostUsersAndItems(ctx context.Context, in *pb.PostUsersAndItemsRequest) (*pb.SimpleApiResponse, error) {
	var req1 pb.PostUsersRequest
	req1.Users = in.Users
	res1 := s.grpcService1PostUsers(ctx, &req1)
	log.Info(res1)

	return &pb.SimpleApiResponse{Status: 200}, nil
}

func (s *MainService) grpcService1Clinet(ctx context.Context, in *empty.Empty) *pb.GetUsersResponse {
	// TODO insecureを直す、inはenptyだとだめなのでinterfaceとかで受ける
	conn, _ := grpc.Dial("grpc-service1-server:4003", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	result, err := c.GetUsers(ctx, in)
	if err != nil {
		log.Error(err)
		return &pb.GetUsersResponse{}
	}

	return result
}

func (s *MainService) grpcService2Clinet(ctx context.Context, in *empty.Empty) *pb.GetItemsResponse {
	// TODO insecureを直す、inはenptyだとだめなのでinterfaceとかで受ける
	conn, _ := grpc.Dial("grpc-service2-server:4004", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewItemServiceClient(conn)

	result, err := c.GetItems(ctx, in)
	if err != nil {
		log.Error(err)
		return &pb.GetItemsResponse{}
	}

	return result
}

func (s *MainService) grpcService1PostUsers(ctx context.Context, in *pb.PostUsersRequest) *pb.SimpleApiResponse {
	conn, _ := grpc.Dial("grpc-service1-server:4003", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	result, err := c.PostUsers(ctx, in)
	if err != nil {
		log.Error(err)
		return &pb.SimpleApiResponse{Status: 500}
	}

	return result
}
