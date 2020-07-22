package models

import (
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
