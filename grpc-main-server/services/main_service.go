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

// // TODO このIDはトランザクション内で共有する値
// var sagaId uint64 = 10

// saga.StorageConfig.Redis.Host = "saga-log-redis-server"
// saga.StorageConfig.Redis.Port = "6379"
// saga.StorageConfig.Redis.Password = "redis"

// tx := saga.AddSubTxDef("test", s.sampleAction, s.sampleCompensate).
// 	CreateSubTx(ctx, sagaId)

// tx.StartSaga().
// 	ExecSub("test", "alice", 100). // 若干クセが有る。メソッドは指定しないが引数は指定するのを直したい
// 	EndSaga()

// func (s *MainService) sampleAction(ctx context.Context, name string, age int) error {
// 	log.Info("action")

// 	if name == "alice" {
// 		return errors.New("owata")
// 	}

// 	return nil
// }

// func (s *MainService) sampleCompensate(ctx context.Context, name string, age int) error {
// 	log.Info("compensate")
// 	log.Info("param: name:" + name + " age:" + strconv.Itoa(age))
// 	return nil
// }
