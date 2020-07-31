package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc/status"
)

type errorBody struct {
	Message string `json:"error,omitempty"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler,
	w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message11"}`
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))

	jsonErr := json.NewEncoder(w).Encode(errorBody{
		Message: err.Error(),
	})

	if jsonErr != nil {
		s, _ := w.Write([]byte(fallback))
		log.Info(s)
	}

	// s, _ := w.Write([]byte(fallback))

	// aaa, ok := errors.Cause(err).(grpcError)

	// log.Info(s)
	// log.Error(aaa)
	// log.Info(ok)
}
