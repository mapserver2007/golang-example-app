package services

import (
	database "github.com/mapserver2007/golang-example-app/openapi-web/common/database"
	"github.com/mapserver2007/golang-example-app/openapi-web/models"
	openapi "github.com/mapserver2007/golang-example-app/openapi-web/openapi/go"
	"gopkg.in/gorp.v1"
)

// ApiService struct
type APIService struct {
	Connection *gorp.DbMap
}

// NewAPIService constructor
func NewAPIService() openapi.ExampleApiServicer {
	return &APIService{Connection: database.GetConnection()}
}

// GetUsers - all users
func (s *APIService) GetUsers() (interface{}, error) {
	db := models.User{Connection: s.Connection}
	rows := db.FindAll()

	users := s.convertUserModelToResponse(rows)
	return openapi.GetUserResponses{
		Users: users,
	}, nil
}

// PostUser - create user
func (s *APIService) PostUser(postUserRequest openapi.PostUserRequest) (interface{}, error) {
	db := models.User{Connection: s.Connection}

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
	db := models.User{Connection: s.Connection}

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
