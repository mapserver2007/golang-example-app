package saga

import (
	"context"

	"gopkg.in/gorp.v1"

	"github.com/mapserver2007/golang-example-app/common/log"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service1-server/models"
)

func CreateUserAction(_ context.Context, conn *gorp.DbMap, users []*pb.PostUserRequest) ([]int64, error) {
	log.Info("createUserAction start")
	var lastInsertIds []int64
	db := models.User{Connection: conn}
	for _, user := range users {
		result, err := db.CreateUser(user)
		if err != nil {
			return lastInsertIds, err
		}
		lastInsertId, _ := result.LastInsertId()
		lastInsertIds = append(lastInsertIds, lastInsertId)
	}
	return lastInsertIds, nil
}

func CreateUserCompensate(_ context.Context, conn *gorp.DbMap, lastInsertIds []int64, users []*pb.PostUserRequest) error {
	log.Info("createUserCompensate start")
	db := models.User{Connection: conn}
	for _, id := range lastInsertIds {
		err := db.CreateUserCompensate(id)
		if err != nil {
			return err
		}
	}

	return nil
}
