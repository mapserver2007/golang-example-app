package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
	services "github.com/mapserver2007/golang-example-app/web/services"
)

func main() {
	log.Printf("Server started")

	APIService := services.NewAPIService()
	ExampleApiController := openapi.NewExampleApiController(APIService)
	router := openapi.NewRouter(ExampleApiController)

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		// handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	)(router)))
}
