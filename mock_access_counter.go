package RateLimiter

import "github.com/stretchr/testify/mock"

type MockAccessCounter struct {
	mock.Mock
}

func (m *MockAccessCounter) Count(ip string, seconds int) int {
	args := m.Called(ip, seconds)
	return args.Int(0)
}
