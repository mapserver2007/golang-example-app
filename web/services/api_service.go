package services

import (
	openapi "github.com/mapserver2007/golang-example-app/web/openapi/go"
)

type ApiService struct{}

func NewApiService() openapi.ExampleApiServicer {
	return &ApiService{}
}

func (s *ApiService) GetAge() (interface{}, error) {
	return openapi.SimpleResponse{
		Result: 200,
	}, nil
}
