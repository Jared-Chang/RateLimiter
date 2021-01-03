package RateLimiter

type AccessCountService struct {
	counter      AccessCounter
	querySeconds int
}

func (s *AccessCountService) QueryByIp(ip string) AccessCount {
	count := s.counter.Count(ip, s.querySeconds)
	return AccessCount{Ip: ip, Count: count}
}

func NewAccessCountService(counter AccessCounter, querySeconds int) *AccessCountService {
	return &AccessCountService{counter: counter, querySeconds: querySeconds}
}

