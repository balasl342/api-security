package middleware

import (
	"context"
	"net/http"
	"time"

	redisCli "github.com/balasl342/api-security/pkg/redis"
	"github.com/go-redis/redis/v8"
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ctx := context.Background()

		count, err := redisCli.Client.Get(ctx, ip).Int()
		if err != nil && err != redis.Nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Rate limit count set to 5. If it exceeds, return with status code 429 ( Too many requests )
		if count >= 5 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		_, err = redisCli.Client.Incr(ctx, ip).Result()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		redisCli.Client.Expire(ctx, ip, time.Minute)

		next.ServeHTTP(w, r)
	})
}
