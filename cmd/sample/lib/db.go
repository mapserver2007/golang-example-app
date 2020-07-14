package lib

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // no problem

	"github.com/go-gorp/gorp"
)

// Database struct
type Database struct {
	connection *gorp.DbMap
}

// Gotbl struct
type Gotbl struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// Connect method
func (database *Database) Connect(dbms string, dbpath string) {
	db, _ := sql.Open(dbms, dbpath)
	database.connection = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}

// Close method
func (database *Database) Close() {
	database.connection.Db.Close()
}

// FindAll method
func (database *Database) FindAll() []Gotbl {
	var gotbl []Gotbl
	database.connection.Select(&gotbl, "select * from gotbl") // TODO エラー処理
	return gotbl
}
