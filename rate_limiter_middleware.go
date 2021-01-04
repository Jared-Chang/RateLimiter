package RateLimiter

import (
	"encoding/json"
	"net/http"
)

type AccessDenied struct {
	Error string `json:"error"`
}

type RateLimiterMiddleware struct {
	accessCounter AccessCounter
}

func NewRateLimiterMiddleware(accessCounter AccessCounter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{accessCounter: accessCounter}
}

func (r *RateLimiterMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	count := r.accessCounter.Count(request.RemoteAddr, 60)

	if count > 60 {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(AccessDenied{"Error"})
	}
}