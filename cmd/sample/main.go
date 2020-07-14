package main

import (
	"fmt"

	lib "github.com/mapserver2007/golang-example-app/cmd/sample/lib"
)

func main() {
	// study1 構造体を使う
	// p1 := &lib.Calc{3, 5}
	// p2 := &lib.Calc2{p1}
	// fmt.Print(p2.Add())

	// study2 クローラを使う
	// c := lib.Crawler{"https://kakaku.com"}
	// c.Crawle("title")

	// study3 DB
	db := lib.Database{}
	db.Connect("mysql", "mariadb:mariadb@tcp(localhost:13340)/godb")
	defer db.Close()

	rows := db.FindAll()
	for _, r := range rows {
		fmt.Println(r.Name)
	}

	// study4 WebServer

	// study4 rest & generate code

	// study5 grpc gateway
}
