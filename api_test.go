package RateLimiter

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ApiTestSuite struct {
	suite.Suite
	TestServer *httptest.Server
}

func TestApiTestSuiteInit(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}

func (t *ApiTestSuite) SetupTest() {
	accessCountHandlerFactor := NewAccessCountHandlerFactor()
	t.TestServer = httptest.NewServer(accessCountHandlerFactor.Create())
}

func (t *ApiTestSuite) TearDownTest() {
	t.TestServer.Close()
}

func (t ApiTestSuite) TestReturnAccessCountByIp() {
	jsonObject := GetResponse(t)

	actual := int(jsonObject["count"].(float64))
	expected := 1

	t.Equal(expected, actual)
}

func (t ApiTestSuite) TestOnlyCanAccess60TimesPerMinute() {
	for i := 0; i < 60; i++ {
		GetResponse(t)
	}
	jsonObject := GetResponse(t)

	actual := jsonObject["error"].(string)
	expected := "Error"

	t.Equal(expected, actual)
}

func (t ApiTestSuite) TestCanAccess1TimePerSecond() {
	for i := 0; i < 120; i++ {
		jsonObject := GetResponse(t)

		fmt.Println("#Access: ", i+1, "Your ip: ", jsonObject["ip"].(string), "#AccessCount: ", int(jsonObject["count"].(float64)))

		actual := int(jsonObject["count"].(float64))

		t.Greater(actual, 0)

		time.Sleep(time.Duration(1100)*time.Millisecond)
	}
}

func GetResponse(t ApiTestSuite) map[string]interface{} {
	response, _ := http.Get(t.TestServer.URL)

	body, _ := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()

	var jsonObject map[string]interface{}
	_ = json.Unmarshal(body, &jsonObject)
	return jsonObject
}
