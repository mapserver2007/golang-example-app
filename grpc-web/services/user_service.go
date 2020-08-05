package services

import (
	"context"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mapserver2007/golang-example-app/grpc-web/common/client"
	"github.com/mapserver2007/golang-example-app/grpc-web/common/log"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-web/models"
	"gopkg.in/gorp.v1"
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

func (s *UserService) PostUser(_ context.Context, in *pb.PostUserRequest) (*pb.SimpleApiResponse, error) {
	db := models.User{Connection: s.Connection}

	if err := db.CreateUser(in); err != nil {
		log.Error(err)
		return &pb.SimpleApiResponse{}, err
	}

	return &pb.SimpleApiResponse{Status: 201}, nil
}

func (s *UserService) PutUser(_ context.Context, in *pb.PutUserRequest) (*pb.SimpleApiResponse, error) {
	db := models.User{Connection: s.Connection}

	result, err := db.UpdateUser(in)
	if err != nil {
		log.Error(err)
		return &pb.SimpleApiResponse{}, err
	}
	rows, _ := result.RowsAffected()
	if rows > 0 {
		return &pb.SimpleApiResponse{Status: 204}, nil
	} else {
		return &pb.SimpleApiResponse{Status: 404}, nil
	}
}

func (s *UserService) GetUsersAndItems(ctx context.Context, in *empty.Empty) (*pb.GetUsersAndItemsRespones, error) {
	result, err := client.GrpcClient("grpc-sub-server:3012", ctx, in)
	if err != nil {
		log.Error(err)
		return &pb.GetUsersAndItemsRespones{}, nil
	}
	items := result.(*pb.GetItemsResponse).Items

	db := models.User{Connection: s.Connection}
	rows, err := db.FindAll()

	if err != nil {
		log.Error(err)
		return &pb.GetUsersAndItemsRespones{}, nil
	}

	var users = []*pb.GetUserResponse{}
	for _, row := range rows {
		users = append(users, &pb.GetUserResponse{Name: row.Name, Age: row.Age})
	}

	return &pb.GetUsersAndItemsRespones{Items: items, Users: users}, nil
}
