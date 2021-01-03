package RateLimiter

import (
	"net/http"
)

type AccessCountHandlerFactor struct {
}

func (f *AccessCountHandlerFactor) Create() http.Handler {
	return NewAccessCountController()
}

func NewAccessCountHandlerFactor() *AccessCountHandlerFactor {
	return &AccessCountHandlerFactor{}
}
