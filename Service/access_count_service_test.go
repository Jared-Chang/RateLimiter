package Service

import (
	"RateLimiter/Model"
	"RateLimiter/TimeSeriesAccessCounter"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccessCountServiceSuite struct {
	suite.Suite
	sut *AccessCountService
	mockAccessCounter *TimeSeriesAccessCounter.MockAccessCounter
}

func TestAccessCountServiceSuiteInit(t *testing.T) {
	suite.Run(t, new(AccessCountServiceSuite))
}

func (t *AccessCountServiceSuite) SetupTest() {
	t.mockAccessCounter = new(TimeSeriesAccessCounter.MockAccessCounter)
	t.sut = NewAccessCountService(t.mockAccessCounter, 60)
}

func (t *AccessCountServiceSuite) TestQueryByIp_60Seconds() {
	ip := "127.0.0.1"
	count := 1

	t.mockAccessCounter.On("Count", ip, 60).Return(1)

	expected := Model.AccessCount{Ip: ip, Count: count}
	actual := t.sut.QueryByIp(ip)

	t.Equal(expected, actual)
}

func (t *AccessCountServiceSuite) TestQueryByIp_Twice_60Seconds() {
	ip := "127.0.0.1"
	count := 2

	t.mockAccessCounter.On("Count", ip, 60).Return(2)

	expected := Model.AccessCount{Ip: ip, Count: count}
	actual := t.sut.QueryByIp(ip)

	t.Equal(expected, actual)
}
