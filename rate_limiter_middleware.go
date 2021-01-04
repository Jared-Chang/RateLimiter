package RateLimiter

import (
	"encoding/json"
	"net/http"
)

type AccessDenied struct {
	Error string `json:"error"`
}

type RateLimiterMiddleware struct {
	AccessCounter AccessCounter
	Handler       http.Handler
}

func NewRateLimiterMiddleware(accessCounter AccessCounter, handler http.Handler) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{AccessCounter: accessCounter, Handler: handler}
}

func (r *RateLimiterMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	count := r.AccessCounter.Count(request.RemoteAddr, 60)

	if count > 60 {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(AccessDenied{"Error"})
	}
}