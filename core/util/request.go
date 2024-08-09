package util

import (
	"github.com/google/uuid"
	"net/http"
)

func GetRequestIdFromHeader(req *http.Request) string {
	requestId := req.Header.Get("x-request-id")
	if requestId == "" {
		requestId = uuid.NewString()
	}
	return requestId
}
