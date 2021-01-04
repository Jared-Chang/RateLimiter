package Controller

import (
	"RateLimiter/Service"
	"RateLimiter/TimeSeriesAccessCounter"
	"encoding/json"
	"net/http"
	"strings"
)

type AccessCountController struct {
	accessCountService *Service.AccessCountService
}

func (a *AccessCountController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response := a.accessCountService.QueryByIp(strings.Split(request.RemoteAddr, ":")[0])
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}

func NewAccessCountController(accessCounter TimeSeriesAccessCounter.AccessCounter, querySeconds int) *AccessCountController {
	return &AccessCountController{accessCountService: Service.NewAccessCountService(accessCounter, querySeconds)}
}
