package services

import (
	models "github.com/mapserver2007/golang-example-app/web/models"
	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
)

// ApiService struct
type APIService struct{}

// NewAPIService constructor
func NewAPIService() openapi.ExampleApiServicer {
	return &APIService{}
}

// GetAge method
func (s *APIService) GetAge() (interface{}, error) {
	db := models.Database{}
	db.Connect("mysql", "mariadb:mariadb@tcp(localhost:13340)/godb")
	defer db.Close()
	users := db.FindAll()
	return openapi.GetUserResponses{
		Users: users,
	}, nil
}
