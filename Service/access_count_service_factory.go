package Service

import (
	"RateLimiter/Controller"
	"RateLimiter/Middleware"
	"RateLimiter/TimeSeriesAccessCounter"
	"net/http"
)

type AccessCountHandlerFactor struct {
}

func (f *AccessCountHandlerFactor) Create() http.Handler {
	return Middleware.NewRateLimiterMiddleware(TimeSeriesAccessCounter.GetInstance(),
		Controller.NewAccessCountController(TimeSeriesAccessCounter.GetInstance(), 60),
		60,
		60)
}

func NewAccessCountHandlerFactor() *AccessCountHandlerFactor {
	return &AccessCountHandlerFactor{}
}
