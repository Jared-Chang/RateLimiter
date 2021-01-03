package RateLimiter

import (
	"RateLimiter/TimeSeriesAccessCounter"
	"encoding/json"
	"net/http"
)

type AccessCountController struct {
	accessCountService *AccessCountService
}

func (a *AccessCountController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response := a.accessCountService.QueryByIp(request.RemoteAddr)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func NewAccessCountController() *AccessCountController {
	return &AccessCountController{accessCountService: NewAccessCountService(TimeSeriesAccessCounter.GetInstance())}
}
