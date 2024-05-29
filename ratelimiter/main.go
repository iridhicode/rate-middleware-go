package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"ratelimiter/middleware"

	"github.com/joho/godotenv"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	capacity, _ := strconv.Atoi(os.Getenv("CAPACITY"))
	leakRateStr := os.Getenv("LEAK_RATE")
	leakRate, _ := time.ParseDuration(leakRateStr)
	whitelistIPs := strings.Split(os.Getenv("WHITELIST_IPS"), ",")
	blacklistIPs := strings.Split(os.Getenv("BLACKLIST_IPS"), ",")

	rateLimitConfig := middleware.RateLimitConfig{
		Capacity:     capacity,
		LeakRate:     leakRate,
		WhitelistIPs: whitelistIPs,
		BlacklistIPs: blacklistIPs,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	handler := middleware.RateLimitMiddleware(rateLimitConfig)(mux)

	log.Printf("Server is running on port %s...", port)
	http.ListenAndServe(":"+port, handler)
}
