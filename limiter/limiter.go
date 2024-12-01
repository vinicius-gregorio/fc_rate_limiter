package limiter

type RateLimiter interface {
	Allow(key string, limit int, duration int64) (bool, error)
	Block(key string, duration int64) error
}
