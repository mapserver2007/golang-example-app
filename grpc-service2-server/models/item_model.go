package models

import (
	"database/sql"

	"gopkg.in/gorp.v1"

	database "github.com/mapserver2007/golang-example-app/common/database"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service2-server/models/sqls"
)

type ItemModel struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

type Item struct {
	Connection *gorp.DbMap
}

func (db *Item) FindAll() ([]ItemModel, error) {
	var result []ItemModel
	_, err := db.Connection.Select(&result, sqls.FindAllItems())
	return result, err
}

func (db *Item) FindById(id int32) (ItemModel, error) {
	var result ItemModel
	err := db.Connection.SelectOne(&result, sqls.FindByItemId(), id)
	return result, err
}

func (db *Item) CreateItem(param *pb.PostItemRequest) (sql.Result, error) {
	result, err := database.TransactionScope(db.Connection, func(tran *gorp.Transaction) (sql.Result, error) {
		return tran.Exec(sqls.CreateItem(), param.Name, param.Price)
	})
	return result, err
}

func (db *Item) CreateItemCompensate(id int64) error {
	_, err := database.TransactionScope(db.Connection, func(tran *gorp.Transaction) (sql.Result, error) {
		return tran.Exec(sqls.CreateItemCompensate(), id)
	})
	return err
}
