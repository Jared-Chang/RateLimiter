package RateLimiter

import (
	"encoding/json"
	"net/http"
)

type AccessCountController struct {
	accessCountService *AccessCountService
}

func (a *AccessCountController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response := a.accessCountService.QueryByIp(request.RemoteAddr)
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}

func NewAccessCountController(accessCounter AccessCounter, querySeconds int) *AccessCountController {
	return &AccessCountController{accessCountService: NewAccessCountService(accessCounter, querySeconds)}
}
