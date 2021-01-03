package RateLimiter

import (
	"net/http"
)

type AccessCountHandlerFactor struct {
}

func (f *AccessCountHandlerFactor) Create() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
	})
}

func NewAccessCountHandlerFactor() *AccessCountHandlerFactor {
	return &AccessCountHandlerFactor{}
}
