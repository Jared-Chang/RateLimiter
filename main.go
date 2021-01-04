package main

import (
	"RateLimiter/ServiceFactory"
	"net/http"
)

func main() {

	accessCountController := ServiceFactory.NewAccessCountHandlerFactor().Create()

	http.Handle("/AccessCount", accessCountController)

	http.ListenAndServe(":8088", nil)
}
