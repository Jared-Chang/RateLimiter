package RateLimiter

import (
	"RateLimiter/TimeSeriesAccessCounter"
	"net/http"
)

type AccessCountHandlerFactor struct {
}

func (f *AccessCountHandlerFactor) Create() http.Handler {
	return NewRateLimiterMiddleware(TimeSeriesAccessCounter.GetInstance(),
		NewAccessCountController(TimeSeriesAccessCounter.GetInstance(), 60),
		60,
		60)
}

func NewAccessCountHandlerFactor() *AccessCountHandlerFactor {
	return &AccessCountHandlerFactor{}
}
