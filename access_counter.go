package RateLimiter

type AccessCounter interface {
	Count(ip string) int
}

