/*
 * golang-example-app
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A ExampleApiController binds http requests to an api service and writes the service results to the http response
type ExampleApiController struct {
	service ExampleApiServicer
}

// NewExampleApiController creates a default api controller
func NewExampleApiController(s ExampleApiServicer) Router {
	return &ExampleApiController{service: s}
}

// Routes returns all of the api route for the ExampleApiController
func (c *ExampleApiController) Routes() Routes {
	return Routes{
		{
			"GetUsers",
			strings.ToUpper("Get"),
			"/v1/users",
			c.GetUsers,
		},
		{
			"PostUser",
			strings.ToUpper("Post"),
			"/v1/users",
			c.PostUser,
		},
		{
			"PutUser",
			strings.ToUpper("Put"),
			"/v1/users/{userId}",
			c.PutUser,
		},
	}
}

// GetUsers - all users
func (c *ExampleApiController) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetUsers()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// PostUser - create user
func (c *ExampleApiController) PostUser(w http.ResponseWriter, r *http.Request) {
	postUserRequest := &PostUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&postUserRequest); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.PostUser(*postUserRequest)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// PutUser - update user
func (c *ExampleApiController) PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	putUserRequest := &PutUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&putUserRequest); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.PutUser(userId, *putUserRequest)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
