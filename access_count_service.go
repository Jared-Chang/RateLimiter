package RateLimiter

type AccessCountService struct {
	counter AccessCounter
}

func (s *AccessCountService) QueryByIp(ip string) AccessCount {
	count := s.counter.Count(ip)
	return AccessCount{Ip: ip, Count: count}
}

func NewAccessCountService(counter AccessCounter) *AccessCountService {
	return &AccessCountService{counter: counter}
}

