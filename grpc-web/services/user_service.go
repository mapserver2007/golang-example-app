package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
)

type UserService struct{}

// func NewUserService() UserService {
// 	return &UserService{}
// }

func (s *UserService) GetUsers(_ context.Context, _ *empty.Empty) (*pb.UsersResponse, error) {
	// TODO
	users := []*pb.User{}
	users = append(users, &pb.User{Name: "test", Age: 10})
	return &pb.UsersResponse{Users: users}, nil
}
