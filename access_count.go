package RateLimiter

type AccessCount struct {
	Ip string `json:"ip"`
	Count int `json:"count"`
}
