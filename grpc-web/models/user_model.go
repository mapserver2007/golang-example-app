package models

import (
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
	database "github.com/mapserver2007/golang-example-app/web/common"
	"gopkg.in/gorp.v1"
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
	_, err := db.Connection.Select(&result, db.sqlFindAll())
	return result, err
}

func (db *User) FindById(id int32) (UserModel, error) {
	var result UserModel
	err := db.Connection.SelectOne(&result, db.sqlFindById(), id)
	return result, err
}

func (db *User) CreateUser(request *pb.PostUserRequest) error {
	db.Connection.AddTableWithName(UserModel{}, "users")
	if err := database.TransactionScope(db.Connection, func(tran *gorp.Transaction) error {
		user := UserModel{Name: request.Name, Age: request.Age}
		return tran.Insert(&user)
	}); err != nil {
		return err
	}
	return nil
}

func (db *User) sqlFindAll() string {
	return `
SELECT
  name,
  age
FROM
  users
`
}

func (db *User) sqlFindById() string {
	return `
SELECT
  name,
  age
FROM
	users
WHERE
	id = ?
`
}
