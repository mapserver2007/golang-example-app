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
	"github.com/mapserver2007/golang-example-app/grpc-service2-server/models"
	"github.com/mapserver2007/golang-example-app/subtx"
)

type ItemService struct {
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

func (s *sagaService) createSubTx(in *pb.PostItemsRequest, conn *gorp.DbMap) {
	subtx.SubTxDefinitions.
		CreateSubTx(s.ctx, conn, s.serverId, in.Uuid).
		ExecSub("createItem", in.Items)
}

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	id64, _ := strconv.ParseInt(in.Id, 10, 64)
	id := int32(id64)

	db := models.Item{Connection: s.Connection}
	row, err := db.FindById(id)

	if err != nil {
		log.Error(err)
		return &pb.GetItemResponse{}, nil
	}

	return &pb.GetItemResponse{Name: row.Name, Price: row.Price}, nil
}

func (s *ItemService) GetItems(ctx context.Context, _ *empty.Empty) (*pb.GetItemsResponse, error) {
	db := models.Item{Connection: s.Connection}
	rows, err := db.FindAll()

	if err != nil {
		log.Error(err)
		return &pb.GetItemsResponse{}, nil
	}

	var items = []*pb.GetItemResponse{}
	for _, row := range rows {
		items = append(items, &pb.GetItemResponse{Name: row.Name, Price: row.Price})
	}
	return &pb.GetItemsResponse{Items: items}, nil
}

func (s *ItemService) PostItems(ctx context.Context, in *pb.PostItemsRequest) (*pb.SimpleApiResponse, error) {
	tx := newSagaService(ctx, "grpc-service2-server")
	tx.createSubTx(in, s.Connection)

	return &pb.SimpleApiResponse{Status: 200}, nil
}
