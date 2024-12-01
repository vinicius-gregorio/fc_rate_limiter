package limiter

import (
	"net/http"

	"github.com/spf13/viper"
)

func NewRateLimiterMiddleware(rl *RedisRateLimiter) func(http.Handler) http.Handler {
	ipLimit := viper.GetInt("RATE_LIMITER_IP_LIMIT")
	tokenLimit := viper.GetInt("RATE_LIMITER_TOKEN_LIMIT")
	blockDuration := viper.GetInt64("RATE_LIMITER_BLOCK_DURATION")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var key string
			token := r.Header.Get("API_KEY")

			if token != "" {
				key = "token:" + token
			} else {
				key = "ip:" + r.RemoteAddr
			}

			limit := ipLimit
			if token != "" {
				limit = tokenLimit
			}

			allowed, err := rl.Allow(key, limit, blockDuration)
			if err != nil || !allowed {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
