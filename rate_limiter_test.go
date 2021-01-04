package RateLimiter

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m.Called(writer, request)
}

type RateLimiterMiddlewareSuite struct {
	suite.Suite
	sut *RateLimiterMiddleware
	mockAccessCounter *MockAccessCounter
	mockHandler *MockHandler
}

func TestRateLimiterMiddlewareSuiteInit(t *testing.T) {
	suite.Run(t, new(RateLimiterMiddlewareSuite))
}

func (t *RateLimiterMiddlewareSuite) SetupTest() {
	t.mockHandler = new(MockHandler)
	t.mockAccessCounter = new(MockAccessCounter)
	t.sut = NewRateLimiterMiddleware(t.mockAccessCounter, t.mockHandler)
}

func (t RateLimiterMiddlewareSuite) TestDeniedAccess_After60TimesAccess_Within1Minute() {

	t.mockAccessCounter.On("Count", mock.Anything, 60).Return(61)

	jsonObject := Get(&t)

	actual := jsonObject["error"].(string)
	expected := "Error"

	t.Equal(expected, actual)
}


func (t RateLimiterMiddlewareSuite) TestAllowAccess() {

	t.mockAccessCounter.On("Count", mock.Anything, 60).Return(1)
	t.mockHandler.On("ServeHTTP", mock.Anything, mock.Anything).Return(nil)

	Get(&t)

	t.mockHandler.AssertCalled(t.T(), "ServeHTTP", mock.Anything, mock.Anything)
}

func Get(t *RateLimiterMiddlewareSuite) (jsonObject map[string]interface{}) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://127.0.0.1", nil)

	t.sut.ServeHTTP(writer, request)

	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	json.Unmarshal(body, &jsonObject)
	return
}