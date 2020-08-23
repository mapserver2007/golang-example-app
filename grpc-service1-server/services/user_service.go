package services

import (
	"context"
	"database/sql"
	"strconv"

	"gopkg.in/gorp.v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mapserver2007/golang-example-app/common/constant"
	"github.com/mapserver2007/golang-example-app/common/log"
	"github.com/mapserver2007/golang-example-app/common/saga"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service1-server/models"
	"github.com/mapserver2007/golang-example-app/subtx"
)

type UserService struct {
	Connection *gorp.DbMap
}

type sagaService struct {
	ctx          context.Context
	serverId     string
	actionResult map[string][]sql.Result
}

func newSagaService(ctx context.Context, serverId string) *sagaService {
	saga.StorageConfig.Redis.Host = constant.SagaLogServerHost
	saga.StorageConfig.Redis.Port = constant.SagaLogServerPort
	saga.StorageConfig.Redis.Password = constant.SagaLogServerPassword

	return &sagaService{
		ctx:          ctx,
		serverId:     serverId,
		actionResult: map[string][]sql.Result{},
	}
}

func (s *sagaService) createSubTx(in *pb.PostUsersRequest, conn *gorp.DbMap) {
	subtx.SubTxDefinitions.
		CreateSubTx(s.ctx, conn, s.serverId, in.Uuid).
		ExecSub("createUser", in.Users)
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
	tx := newSagaService(ctx, "grpc-service1-server")
	tx.createSubTx(in, s.Connection)

	return &pb.SimpleApiResponse{Status: 200}, nil
}
