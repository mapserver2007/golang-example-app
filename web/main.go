package main

import (
	"log"
	"net/http"

	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
	services "github.com/mapserver2007/golang-example-app/web/services"
)

func main() {
	log.Printf("Server started")

	APIService := services.NewAPIService()
	ExampleApiController := openapi.NewExampleApiController(APIService)

	router := openapi.NewRouter(ExampleApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
