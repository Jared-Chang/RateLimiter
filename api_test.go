package RateLimiter

import (
	"RateLimiter/Service"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiTestSuite struct {
	suite.Suite
	TestServer *httptest.Server
}

func TestApiTestSuiteInit(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}

func (t *ApiTestSuite) SetupTest() {
	accessCountHandlerFactor := Service.NewAccessCountHandlerFactor()
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

func GetResponse(t ApiTestSuite) map[string]interface{} {
	response, _ := http.Get(t.TestServer.URL)

	body, _ := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()

	var jsonObject map[string]interface{}
	_ = json.Unmarshal(body, &jsonObject)
	return jsonObject
}
