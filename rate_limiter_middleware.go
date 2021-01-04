package RateLimiter

import (
	"net/http"
)

type RateLimiterMiddleware struct {
	accessCounter AccessCounter
}

func NewRateLimiterMiddleware(accessCounter AccessCounter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{accessCounter: accessCounter}
}

func (r *RateLimiterMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}