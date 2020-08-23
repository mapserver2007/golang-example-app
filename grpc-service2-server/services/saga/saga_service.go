package saga

import (
	"context"

	"gopkg.in/gorp.v1"

	"github.com/mapserver2007/golang-example-app/common/log"
	_ "github.com/mapserver2007/golang-example-app/common/saga/storage/redis"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service2-server/models"
)

func CreateItemAction(_ context.Context, conn *gorp.DbMap, items []*pb.PostItemRequest) ([]int64, error) {
	log.Info("createItemAction start")
	var lastInsertIds []int64
	db := models.Item{Connection: conn}
	for _, item := range items {
		result, err := db.CreateItem(item)
		if err != nil {
			return lastInsertIds, err
		}
		lastInsertId, _ := result.LastInsertId()
		lastInsertIds = append(lastInsertIds, lastInsertId)
	}
	return lastInsertIds, nil
}

func CreateItemCompensate(_ context.Context, conn *gorp.DbMap, lastInsertIds []int64, items []*pb.PostItemRequest) error {
	log.Info("createItemCompensate start")
	db := models.Item{Connection: conn}
	for _, id := range lastInsertIds {
		err := db.CreateItemCompensate(id)
		if err != nil {
			return err
		}
	}
	return nil
}
