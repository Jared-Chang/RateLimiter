package RateLimiter

type AccessCountService struct {
}

func (s *AccessCountService) QueryByIp(ip string) AccessCount {
	return AccessCount{}
}

func NewAccessCountService() *AccessCountService {
	return &AccessCountService{}
}

