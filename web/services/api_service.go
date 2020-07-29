package services

import (
	"github.com/mapserver2007/golang-example-app/web/common/database"
	"github.com/mapserver2007/golang-example-app/web/models"
	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
)

// ApiService struct
type APIService struct{}

// NewAPIService constructor
func NewAPIService() openapi.ExampleApiServicer {
	return &APIService{}
}

// GetUsers - all users
func (s *APIService) GetUsers() (interface{}, error) {
	conn := database.GetConnection()
	defer conn.Db.Close()

	db := models.User{Connection: conn}
	rows := db.FindAll()

	users := s.convertUserModelToResponse(rows)
	return openapi.GetUserResponses{
		Users: users,
	}, nil
}

// PostUser - create user
func (s *APIService) PostUser(postUserRequest openapi.PostUserRequest) (interface{}, error) {
	conn := database.GetConnection()
	defer conn.Db.Close()

	db := models.User{Connection: conn}

	if err := db.CreateUser(postUserRequest); err != nil {
		return openapi.SimpleStatusResponse{
			Status: 500,
		}, err
	}

	return openapi.SimpleStatusResponse{
		Status: 201,
	}, nil
}

// PutUser - update user
func (s *APIService) PutUser(userId string, putUserRequest openapi.PutUserRequest) (interface{}, error) {
	conn := database.GetConnection()
	defer conn.Db.Close()

	db := models.User{Connection: conn}

	if err := db.UpdateUser(userId, putUserRequest); err != nil {
		return openapi.SimpleStatusResponse{
			Status: 500,
		}, err
	}

	return openapi.SimpleStatusResponse{
		Status: 204,
	}, nil
}

func (s *APIService) convertUserModelToResponse(list []models.UserModel) []openapi.GetUserResponse {
	var responses = []openapi.GetUserResponse{}
	for _, elem := range list {
		responses = append(responses, openapi.GetUserResponse{Name: elem.Name, Age: elem.Age})
	}
	return responses
}
