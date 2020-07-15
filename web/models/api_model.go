package models

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
)

type Database struct {
	connection *gorp.DbMap
}

// Connect method
func (database *Database) Connect(dbms, dbpath string) {
	db, err := sql.Open(dbms, dbpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	database.connection = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
}

// Close method
func (database *Database) Close() {
	database.connection.Db.Close()
}

// FindAll method
func (database *Database) FindAll() []openapi.GetUserResponse {
	var result []openapi.GetUserResponse
	_, err := database.connection.Select(&result, "select name, age from gotbl")
	if err != nil {
		log.Fatal(err)
	}
	return result
}
