package services

import (
	"context"
	"strconv"

	"gopkg.in/gorp.v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mapserver2007/golang-example-app/common/log"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service1-server/models"
)

type UserService struct {
	Connection *gorp.DbMap
}

func (s *UserService) GetUser(_ context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id64, _ := strconv.ParseInt(in.Id, 10, 64)
	id := int32(id64)

	db := models.User{Connection: s.Connection}
	row, err := db.FindById(id)

	if err != nil {
		log.Error(err)
		return &pb.GetUserResponse{}, nil
	}

	return &pb.GetUserResponse{Name: row.Name, Age: row.Age}, nil
}

func (s *UserService) GetUsers(ctx context.Context, in *empty.Empty) (*pb.GetUsersResponse, error) {
	db := models.User{Connection: s.Connection}
	rows, err := db.FindAll()

	if err != nil {
		log.Error(err)
		return &pb.GetUsersResponse{}, nil
	}

	var users = []*pb.GetUserResponse{}
	for _, row := range rows {
		users = append(users, &pb.GetUserResponse{Name: row.Name, Age: row.Age})
	}
	return &pb.GetUsersResponse{Users: users}, nil
}

func (s *UserService) PostUsers(ctx context.Context, in *pb.PostUsersRequest) (*pb.SimpleApiResponse, error) {
	tx := newSagaService(ctx, "grpc-service1-server", s.Connection)
	tx.createUserSubTx(in)

	return &pb.SimpleApiResponse{Status: 200}, nil
}
