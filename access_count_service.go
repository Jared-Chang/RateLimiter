package RateLimiter

type AccessCountService struct {
}

func (s *AccessCountService) QueryByIp(ip string) AccessCount {
	return AccessCount{Ip: ip, Count: 1}
}

func NewAccessCountService() *AccessCountService {
	return &AccessCountService{}
}

