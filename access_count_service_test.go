package RateLimiter

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockAccessCounter struct {
	mock.Mock
}

func (m *MockAccessCounter) Count(ip string) int {
	args := m.Called(ip)
	return args.Int(0)
}

type AccessCountServiceSuite struct {
	suite.Suite
	sut *AccessCountService
	mockAccessCounter *MockAccessCounter
}

func TestAccessCountServiceSuiteInit(t *testing.T) {
	suite.Run(t, new(AccessCountServiceSuite))
}

func (t *AccessCountServiceSuite) SetupTest() {
	t.mockAccessCounter = new(MockAccessCounter)
	t.sut = NewAccessCountService(t.mockAccessCounter)
}

func (t *AccessCountServiceSuite) TestQueryByIp() {
	ip := "127.0.0.1"
	count := 1

	t.mockAccessCounter.On("Count", ip).Return(1)

	expected := AccessCount{Ip: ip, Count: count}
	actual := t.sut.QueryByIp(ip)

	t.Equal(expected, actual)
}

func (t *AccessCountServiceSuite) TestQueryByIp_Twice() {
	ip := "127.0.0.1"
	count := 2

	t.mockAccessCounter.On("Count", ip).Return(2)

	expected := AccessCount{Ip: ip, Count: count}
	actual := t.sut.QueryByIp(ip)

	t.Equal(expected, actual)
}
