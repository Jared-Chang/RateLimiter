package RateLimiter

type AccessCounter interface {
	Count(ip string, seconds int) int
	Insert(ip string)
}

