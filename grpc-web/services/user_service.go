package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
)

type UserService struct{}

func (s *UserService) GetUser(_ context.Context, _ *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Name: "test", Age: 20}, nil
}

func (s *UserService) GetUsers(_ context.Context, _ *empty.Empty) (*pb.GetUsersResponse, error) {
	users := []*pb.GetUserResponse{}
	users = append(users, &pb.GetUserResponse{Name: "test", Age: 10})
	return &pb.GetUsersResponse{Users: users}, nil
}
