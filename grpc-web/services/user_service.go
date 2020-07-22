package services

import (
	"context"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	database "github.com/mapserver2007/golang-example-app/grpc-web/common/database"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-web/models"
)

type UserService struct{}

func (s *UserService) GetUser(_ context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id64, _ := strconv.ParseInt(in.Id, 10, 64)
	id := int32(id64)

	conn := database.GetConnection()
	defer conn.Db.Close()

	db := models.User{Connection: conn}
	row, err := db.FindById(id)

	if err != nil {
		return &pb.GetUserResponse{}, err
	}

	return &pb.GetUserResponse{Name: row.Name, Age: row.Age}, nil
}

func (s *UserService) GetUsers(_ context.Context, _ *empty.Empty) (*pb.GetUsersResponse, error) {
	conn := database.GetConnection()
	defer conn.Db.Close()

	db := models.User{Connection: conn}
	rows, err := db.FindAll()

	if err != nil {
		return &pb.GetUsersResponse{}, nil
	}

	var users = []*pb.GetUserResponse{}
	for _, row := range rows {
		users = append(users, &pb.GetUserResponse{Name: row.Name, Age: row.Age})
	}
	return &pb.GetUsersResponse{Users: users}, nil
}

func (s *UserService) PostUser(_ context.Context, in *pb.PostUserRequest) (*pb.SimpleApiResponse, error) {
	// TODO
	return &pb.SimpleApiResponse{Status: 204}, nil
}
