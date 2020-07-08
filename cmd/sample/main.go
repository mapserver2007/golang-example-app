package main

import (
	"github.com/mapserver2007/golang-example-app/cmd/sample/lib"
)

func main() {
	// study1 構造体を使う
	// p1 := &lib.Calc{3, 5}
	// p2 := &lib.Calc2{p1}
	// fmt.Print(p2.Add())

	// study2 クローラを使う
	c := lib.Crawler{"https://kakaku.com"}
	c.Crawle("title")
}
