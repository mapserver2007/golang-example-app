package http

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/mapserver2007/golang-example-app/common/log"
)

type errorBody struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler,
	w http.ResponseWriter, _ *http.Request, err error) {
	statusCode := runtime.HTTPStatusFromCode(status.Code(err))
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(statusCode)

	if json.NewEncoder(w).Encode(errorBody{
		Message: status.Convert(err).Message(),
		Status:  statusCode,
	}) != nil {
		log.Fatal("Response error")
	}
}
