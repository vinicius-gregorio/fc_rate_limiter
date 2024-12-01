package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/vinicius-gregorio/fc_rate_limiter/limiter"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Could not load .env file: %v", err)
	}

	redisAddr := viper.GetString("REDIS_HOST") + ":" + viper.GetString("REDIS_PORT")
	redisPassword := viper.GetString("REDIS_PASSWORD")

	rateLimiter := limiter.NewRedisRateLimiter(redisAddr, redisPassword)

	rateLimiterMiddleware := limiter.NewRateLimiterMiddleware(rateLimiter)

	/// Middleware send injetado na rota
	http.Handle("/", rateLimiterMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})))

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
