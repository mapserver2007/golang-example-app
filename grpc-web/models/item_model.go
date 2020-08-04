package models

import (

	"gopkg.in/gorp.v1"
	"github.com/mapserver2007/golang-example-app/grpc-web/models/sqls"
)

type ItemModel struct {
	Name string `json:"name"`
	Price int32 `json:"price"`
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
