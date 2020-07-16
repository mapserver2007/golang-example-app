package models

import (
	"log"

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

func (db *User) CreateUser(request openapi.PostUserRequest) error {
	db.Connection.AddTableWithName(UserModel{}, "users")
	user := UserModel{Name: request.Name, Age: request.Age}
	tran, err := db.Connection.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = tran.Insert(&user)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return tran.Commit()
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
