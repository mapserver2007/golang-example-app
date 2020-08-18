package services

import (
	"context"
	"strconv"

	"gopkg.in/gorp.v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mapserver2007/golang-example-app/common/log"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service2-server/models"
)

type ItemService struct {
	Connection *gorp.DbMap
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
	tx := newSagaService(ctx, "grpc-service2-server", s.Connection)
	tx.createItemSubTx(in)

	return &pb.SimpleApiResponse{Status: 200}, nil
}
