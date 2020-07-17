package models

import (
	"log"

	database "github.com/mapserver2007/golang-example-app/web/common"
	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
	"gopkg.in/gorp.v1"
)

type UserModel struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type User struct {
	Connection *gorp.DbMap
}

func (db *User) FindAll() []UserModel {
	var result []UserModel
	_, err := db.Connection.Select(&result, db.sqlFindAll())
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *User) CreateUser(request openapi.PostUserRequest) (err error) {
	if err = database.TransactionScope(db.Connection, func(tran *gorp.Transaction) error {
		db.Connection.AddTableWithName(UserModel{}, "users")
		user := UserModel{Name: request.Name, Age: request.Age}
		return tran.Insert(&user)
	}); err != nil {
		log.Fatal(err)
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
