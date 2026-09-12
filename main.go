package main

import (
	"fmt"
	"time"

	"go-ddos-guard/ratelimiter"
)

func main() {
	rL := ratelimiter.NewRateLimiter()

	for i := 1; i <= 6; i++ {
		entry := ratelimiter.LogEntry{ClientID: "client-1", Timestamp: time.Now()}
		allowed := rL.Allow(entry)
		fmt.Printf("request %d: allowed=%t\n", i, allowed)
	}
}
