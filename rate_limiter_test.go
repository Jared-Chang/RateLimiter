package RateLimiter

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type RateLimiterMiddlewareSuite struct {
	suite.Suite
	sut *RateLimiterMiddleware
	mockAccessCounter *MockAccessCounter
}

func TestRateLimiterMiddlewareSuiteInit(t *testing.T) {
	suite.Run(t, new(RateLimiterMiddlewareSuite))
}

func (t *RateLimiterMiddlewareSuite) SetupTest() {
	t.mockAccessCounter = new(MockAccessCounter)
	t.sut = NewRateLimiterMiddleware(t.mockAccessCounter)
}

func (t RateLimiterMiddlewareSuite) TestDeniedAccess_After60TimesAccess_Within1Minute() {

	t.mockAccessCounter.On("Count", mock.AnythingOfType("string"), 60).Return(61)

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://127.0.0.1", nil)

	t.sut.ServeHTTP(writer, request)

	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	fmt.Println(string(body))

	var jsonObject map[string]interface{}
	json.Unmarshal(body, &jsonObject)

	actual := jsonObject["error"].(string)
	expected := "Error"

	t.Equal(expected, actual)
}
