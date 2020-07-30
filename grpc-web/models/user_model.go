package models

import (
	"database/sql"

	database "github.com/mapserver2007/golang-example-app/grpc-web/common/database"
	pb "github.com/mapserver2007/golang-example-app/grpc-web/gen/go"
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

func (db *User) CreateUser(request *pb.PostUserRequest) (err error) {
	db.Connection.AddTableWithName(UserModel{}, "users")
	_, err = database.TransactionScope(db.Connection, func(tran *gorp.Transaction) (sql.Result, error) {
		user := UserModel{Name: request.Name, Age: request.Age}
		return nil, tran.Insert(&user)
	})
	return
}

func (db *User) UpdateUser(request *pb.PutUserRequest) (result sql.Result, err error) {
	result, err = database.TransactionScope(db.Connection, func(tran *gorp.Transaction) (sql.Result, error) {
		return tran.Exec(db.sqlUpdateUserById(), request.Name, request.Age, request.Id)
	})
	return
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

func (db *User) sqlUpdateUserById() string {
	return `
UPDATE
	users
SET
	name = ?,
	age = ?
WHERE
	id = ?
`
}
