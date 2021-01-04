package Middleware

import (
	"RateLimiter/TimeSeriesAccessCounter"
	"encoding/json"
	"net/http"
	"sync"
)

type AccessDenied struct {
	Error string `json:"error"`
}

type RateLimiterMiddleware struct {
	AccessCounter TimeSeriesAccessCounter.AccessCounter
	Handler       http.Handler
	Seconds       int
	LimitCount    int
}

func NewRateLimiterMiddleware(accessCounter TimeSeriesAccessCounter.AccessCounter, handler http.Handler, seconds int, limitCount int) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{AccessCounter: accessCounter, Handler: handler, Seconds: seconds, LimitCount: limitCount}
}

func (r *RateLimiterMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	m := new(sync.Mutex)
	m.Lock()
	r.AccessCounter.Insert(request.RemoteAddr)
	count := r.AccessCounter.Count(request.RemoteAddr, r.Seconds)
	m.Unlock()

	if count > r.LimitCount {
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(AccessDenied{"Error"})
		return
	}

	r.Handler.ServeHTTP(writer, request)
}