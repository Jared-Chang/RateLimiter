package RateLimiter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccessCountServiceSuite struct {
	suite.Suite
	*AccessCountService
}

func TestAccessCountServiceSuiteInit(t *testing.T) {
	suite.Run(t, new(AccessCountServiceSuite))
}

func (t *AccessCountServiceSuite) SetupTest() {
	t.AccessCountService = NewAccessCountService()
}

func (t AccessCountServiceSuite) TestQueryByIp() {
	ip := "127.0.0.1"
	count := 1

	expected := AccessCount{Ip: ip, Count: count}
	actual := t.AccessCountService.QueryByIp(ip)

	t.Equal(expected, actual)
}

func (t AccessCountServiceSuite) TestQueryByIp_Twice() {
	ip := "127.0.0.1"
	count := 2

	expected := AccessCount{Ip: ip, Count: count}
	actual := t.AccessCountService.QueryByIp(ip)

	t.Equal(expected, actual)
}
