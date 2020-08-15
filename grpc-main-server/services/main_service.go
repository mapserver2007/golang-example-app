package services

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mapserver2007/golang-example-app/common/log"
	"github.com/mapserver2007/golang-example-app/common/saga"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
)

type MainService struct {
}

func (s *MainService) GetUsersAndItems(ctx context.Context, in *empty.Empty) (*pb.GetUsersAndItemsResponse, error) {
	s.execSaga(ctx)
	users := s.grpcService1Clinet(ctx, in).Users
	items := s.grpcService2Clinet(ctx, in).Items

	return &pb.GetUsersAndItemsResponse{Users: users, Items: items}, nil
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

func (s *MainService) execSaga(ctx context.Context) {
	var sagaId uint64 = 10

	saga.StorageConfig.Redis.Host = "saga-log-redis-server"
	saga.StorageConfig.Redis.Port = "6379"
	saga.StorageConfig.Redis.Password = "redis"

	tx := saga.AddSubTxDef("test", sampleAction, sampleCompensate).
		InitSaga(ctx, sagaId)

	tx.StartSaga().
		ExecSub("test", "alice", 100). // 若干クセが有る。メソッドは指定しないが引数は指定するのを直したい
		EndSaga()
}

func sampleAction(ctx context.Context, name string, age int) error {
	log.Info("action")

	if name == "alice" {
		return errors.New("owata")
	}

	return nil
}

func sampleCompensate(ctx context.Context, name string, age int) error {
	log.Info("compensate")
	log.Info("param: name:" + name + " age:" + strconv.Itoa(age))
	return nil
}
