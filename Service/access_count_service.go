package Service

import (
	"RateLimiter/Model"
	"RateLimiter/TimeSeriesAccessCounter"
)

type AccessCountService struct {
	counter      TimeSeriesAccessCounter.AccessCounter
	querySeconds int
}

func (s *AccessCountService) QueryByIp(ip string) Model.AccessCount {
	count := s.counter.Count(ip, s.querySeconds)
	return Model.AccessCount{Ip: ip, Count: count}
}

func NewAccessCountService(counter TimeSeriesAccessCounter.AccessCounter, querySeconds int) *AccessCountService {
	return &AccessCountService{counter: counter, querySeconds: querySeconds}
}

