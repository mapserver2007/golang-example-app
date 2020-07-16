package models

import (
	"log"

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

func (db *User) sqlFindAll() string {
	return `
SELECT
  name,
  age
FROM
  users
`
}
