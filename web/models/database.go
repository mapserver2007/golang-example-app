package models

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

func GetConnection(dbms, userId, password, host, port, database string) *gorp.DbMap {
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
