package services

import (
	"context"

	"github.com/mapserver2007/golang-example-app/common/constant"
	"github.com/mapserver2007/golang-example-app/common/log"
	"github.com/mapserver2007/golang-example-app/common/saga"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
)

type sagaService struct {
	ctx      context.Context
	serverId string
}

func newSagaService(ctx context.Context, serverId string) *sagaService {
	saga.StorageConfig.Redis.Host = constant.SagaLogServerHost
	saga.StorageConfig.Redis.Port = constant.SagaLogServerPort
	saga.StorageConfig.Redis.Password = constant.SagaLogServerPassword

	return &sagaService{
		ctx:      ctx,
		serverId: serverId,
	}
}

func (s *sagaService) subTx() {
	var sagaId uint64 = 10
	saga.AddSubTxDef("test", s.action1, s.compensate1).
		CreateSubTx(s.ctx, s.serverId, sagaId).
		ExecSub("test", "alice")
}

func (s *sagaService) action1(_ context.Context, name string) {
	log.Info("action1: " + name)
}

func (s *sagaService) compensate1(_ context.Context, name string) {
	log.Info("compensate1: " + name)
}
