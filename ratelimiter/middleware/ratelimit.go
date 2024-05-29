package middleware

import (
	"net/http"
	"strings"
	"time"

	"ratelimiter/limiter"
)

type RateLimitConfig struct {
	Capacity     int
	LeakRate     time.Duration
	WhitelistIPs []string
	BlacklistIPs []string
}

func RateLimitMiddleware(config RateLimitConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		limiter := limiter.NewLeakyBucketLimiter(config.Capacity, config.LeakRate)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := getClientIP(r)

			if isWhitelisted(clientIP, config.WhitelistIPs) {
				next.ServeHTTP(w, r)
				return
			}

			if isBlacklisted(clientIP, config.BlacklistIPs) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			if !limiter.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getClientIP(r *http.Request) string {
	// Implement logic to extract the client IP from the request
	// This can be done using the X-Forwarded-For header or other methods
	// depending on your server setup and proxy configuration
	// For simplicity, let's assume the client IP is directly accessible
	return r.RemoteAddr
}

func isWhitelisted(clientIP string, whitelist []string) bool {
	for _, ip := range whitelist {
		if ip == clientIP {
			return true
		}
	}
	return false
}

func isBlacklisted(clientIP string, blacklist []string) bool {
	for _, ip := range blacklist {
		if strings.HasPrefix(clientIP, ip) {
			return true
		}
	}
	return false
}
