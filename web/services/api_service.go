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
	conn := models.GetConnection("mysql", "mariadb", "mariadb", "localhost", "13340", "godb")
	defer conn.Db.Close()

	db := models.User{Connection: conn}
	rows := db.FindAll()

	users := s.convertUserModelToResponse(rows)
	return openapi.GetUserResponses{
		Users: users,
	}, nil
}

func (s *APIService) convertUserModelToResponse(list []models.UserModel) []openapi.GetUserResponse {
	var responses = []openapi.GetUserResponse{}
	for _, elem := range list {
		responses = append(responses, openapi.GetUserResponse{Name: elem.Name, Age: elem.Age})
	}
	return responses
}
