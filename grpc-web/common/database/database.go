package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	constant "github.com/mapserver2007/golang-example-app/grpc-web/common/constant"
	"gopkg.in/gorp.v1"
)

type DB struct {
	DBMS     string `yaml:"dbms"`
	UserId   string `yaml:"user_id"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func GetConnection() *gorp.DbMap {
	db, err := sql.Open(
		constant.DBMS,
		fmt.Sprint(constant.DBUserId, ":", constant.DBPassword, "@(", constant.DBHost, ":", constant.DBPort, ")/", constant.DBName),
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
}

// Transaction Scope - transaction wrapper
func TransactionScope(db *gorp.DbMap, tranFunc func(*gorp.Transaction) error) (err error) {
	var tran *gorp.Transaction
	tran, err = db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer func() {
		if proc := recover(); proc != nil {
			_ = tran.Rollback()
			panic(proc)
		} else if err != nil {
			_ = tran.Rollback()
		} else {
			_ = tran.Commit()
			err = nil
		}
	}()

	err = tranFunc(tran)

	return
}
