package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"gopkg.in/yaml.v2"
)

type DB struct {
	DBMS     string `yaml:"dbms"`
	UserId   string `yaml:"user_id"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

var configPath = "web/conf/db.yml"

func GetConnection() *gorp.DbMap {
	dbconfig, err := loadConfig()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	db, err := sql.Open(
		dbconfig.DBMS,
		fmt.Sprint(dbconfig.UserId, ":", dbconfig.Password, "@(", dbconfig.Host, ":", dbconfig.Port, ")/", dbconfig.DBName),
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

func loadConfig() (DB, error) {
	var dbconfig DB
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
		return dbconfig, err
	}
	if err = yaml.UnmarshalStrict(buf, &dbconfig); err != nil {
		log.Fatal(err)
		return dbconfig, err
	}

	return dbconfig, nil
}
