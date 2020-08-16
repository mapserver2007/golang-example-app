package services

import (
	"context"

	"github.com/mapserver2007/golang-example-app/common/constant"
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

func (s *sagaService) start() {
	// TODO sagaIdはstart時RDBでキーを発行するようにする
	var sagaId uint64 = 10
	saga.CreateSubTx(s.ctx, s.serverId, sagaId).StartSaga()
}

func (s *sagaService) end() {
	var sagaId uint64 = 10
	saga.CreateSubTx(s.ctx, s.serverId, sagaId).EndSaga()
}
