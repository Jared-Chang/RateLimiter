package RateLimiter

import (
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
	accessCountHandlerFactor := NewAccessCountHandlerFactor()
	t.TestServer = httptest.NewServer(accessCountHandlerFactor.Create())
}

func (t *ApiTestSuite) TearDownTest() {
	t.TestServer.Close()
}

func (t ApiTestSuite) TestReturnAccessCountByIp() {
	response, _ := http.Get(t.TestServer.URL)

	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var jsonObject map[string]interface{}
	json.Unmarshal(body, &jsonObject)

	actual := int(jsonObject["count"].(float64))
	expected := 1

	t.Equal(expected, actual)
}