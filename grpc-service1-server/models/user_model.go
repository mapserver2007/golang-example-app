package models

import (
	"database/sql"

	"gopkg.in/gorp.v1"

	database "github.com/mapserver2007/golang-example-app/common/database"
	pb "github.com/mapserver2007/golang-example-app/gen/go"
	"github.com/mapserver2007/golang-example-app/grpc-service1-server/models/sqls"
)

type UserModel struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type User struct {
	Connection *gorp.DbMap
}

func (db *User) FindAll() ([]UserModel, error) {
	var result []UserModel
	_, err := db.Connection.Select(&result, sqls.FindAllUsers())
	return result, err
}

func (db *User) FindById(id int32) (UserModel, error) {
	var result UserModel
	err := db.Connection.SelectOne(&result, sqls.FindByUserId(), id)
	return result, err
}

func (db *User) CreateUser(param *pb.PostUserRequest) (sql.Result, error) {
	db.Connection.AddTableWithName(UserModel{}, "users")
	result, err := database.TransactionScope(db.Connection, func(tran *gorp.Transaction) (sql.Result, error) {
		return tran.Exec(sqls.CreateUser(), param.Name, param.Age)
	})
	return result, err
}

func (db *User) CreateUserCompensate(id int64) error {
	_, err := database.TransactionScope(db.Connection, func(tran *gorp.Transaction) error {
		return tran.Exec(sqls.CreateUserCompensate(), id)
	})
	return err
}

func (db *User) UpdateUser(param *pb.PutUserRequest) (sql.Result, error) {
	result, err = database.TransactionScope(db.Connection, func(tran *gorp.Transaction) (sql.Result, error) {
		return tran.Exec(sqls.UpdateByUserId(), param.Name, param.Age, param.Id)
	})
	return result, err
}
