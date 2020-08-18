package services

import (
	"context"
	"database/sql"

	"gopkg.in/gorp.v1"

	"github.com/mapserver2007/golang-example-app/common/constant"
	"github.com/mapserver2007/golang-example-app/common/log"
	"github.com/mapserver2007/golang-example-app/common/saga"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service2-server/models"
)

type sagaService struct {
	ctx          context.Context
	serverId     string
	conn         *gorp.DbMap
	actionResult map[string][]sql.Result
}

func newSagaService(ctx context.Context, serverId string, conn *gorp.DbMap) *sagaService {
	saga.StorageConfig.Redis.Host = constant.SagaLogServerHost
	saga.StorageConfig.Redis.Port = constant.SagaLogServerPort
	saga.StorageConfig.Redis.Password = constant.SagaLogServerPassword

	return &sagaService{
		ctx:          ctx,
		serverId:     serverId,
		conn:         conn,
		actionResult: map[string][]sql.Result{},
	}
}

func (s *sagaService) createItemSubTx(in *pb.PostItemsRequest) {
	// TODO メインサーバで発行、headerで引き回したい
	var sagaId uint64 = 10
	saga.AddSubTxDef("createItem", s.createItemAction, s.createItemCompensate).
		CreateSubTx(s.ctx, s.serverId, sagaId).
		ExecSub("createItem", in.Items)
}

func (s *sagaService) createItemAction(_ context.Context, items []*pb.PostItemRequest) error {
	log.Info("createItemAction start")
	db := models.Item{Connection: s.conn}
	for _, item := range items {
		result, err := db.CreateItem(item)
		if err != nil {
			return err
		}
		s.actionResult["createItemAction"] = append(s.actionResult["createItemAction"], result)
	}

	return nil
}

func (s *sagaService) createItemCompensate(_ context.Context, items []*pb.PostItemRequest) error {
	log.Info("createItemCompensate start")
	db := models.Item{Connection: s.conn}
	for idx := 0; idx < len(items); idx++ {
		id, _ := s.actionResult["createItemAction"][idx].LastInsertId()
		err := db.CreateItemCompensate(id)
		if err != nil {
			return err
		}
	}

	return nil
}
