package services

import (
	"context"
	"database/sql"
	"errors"

	"gopkg.in/gorp.v1"

	"github.com/mapserver2007/golang-example-app/common/constant"
	"github.com/mapserver2007/golang-example-app/common/log"
	"github.com/mapserver2007/golang-example-app/common/saga"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service1-server/models"
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

func (s *sagaService) createUserSubTx(in *pb.PostUsersRequest) {
	// TODO メインサーバで発行、headerで引き回したい
	var sagaId uint64 = 10
	saga.AddSubTxDef("createUser", s.createUserAction, s.createUserCompensate).
		CreateSubTx(s.ctx, s.serverId, sagaId).
		ExecSub("createUser", in.Users)
}

func (s *sagaService) createUserAction(_ context.Context, users []*pb.PostUserRequest) error {
	log.Info("action start")
	db := models.User{Connection: s.conn}
	for _, user := range users {
		result, err := db.CreateUser(user)
		if err != nil {
			return err
		}
		s.actionResult["createUserAction"] = append(s.actionResult["createUserAction"], result)
	}

	return errors.New("owata")
}

func (s *sagaService) createUserCompensate(_ context.Context, users []*pb.PostUserRequest) error {
	log.Info("compensate start")
	db := models.User{Connection: s.conn}
	for idx := 0; idx < len(users); idx++ {
		id, _ := s.actionResult["createUserAction"][idx].LastInsertId()
		err := db.CreateUserCompensate(id)
		if err != nil {
			return err
		}
	}

	return nil
}
