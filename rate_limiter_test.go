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
	t.sut = NewRateLimiterMiddleware(t.mockAccessCounter, t.mockHandler, 60, 60)
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

func (t RateLimiterMiddlewareSuite) TestDeniedAccess_After30TimesAccess_Within30Seconds() {

	t.sut = NewRateLimiterMiddleware(t.mockAccessCounter, t.mockHandler, 30, 30)
	t.mockAccessCounter.On("Count", mock.Anything, 30).Return(31)

	jsonObject := Get(&t)

	actual := jsonObject["error"].(string)
	expected := "Error"

	t.Equal(expected, actual)
}

func (t RateLimiterMiddlewareSuite) TestIncreaseAccessCount() {

	remoteAddr := "1.2.3.4"

	t.mockAccessCounter.On("Count", mock.Anything, mock.Anything).Return(61)
	t.mockAccessCounter.On("Insert", remoteAddr).Return(nil)

	_ = GetWithRemoteAddr(&t, remoteAddr)

	t.mockAccessCounter.AssertCalled(t.T(), "Insert", remoteAddr)
}

func GetWithRemoteAddr(t *RateLimiterMiddlewareSuite, remoteAddr string) (jsonObject map[string]interface{}) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://127.0.0.1", nil)
	request.RemoteAddr = remoteAddr

	t.sut.ServeHTTP(writer, request)

	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()

	_ = json.Unmarshal(body, &jsonObject)
	return
}

func Get(t *RateLimiterMiddlewareSuite) (jsonObject map[string]interface{}) {
	return GetWithRemoteAddr(t, "127.0.0.1")
}