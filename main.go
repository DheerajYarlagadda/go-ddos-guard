package main

import (
	"fmt"
	"time"
)

func main() {
	rL := NewRateLimiter()
	entry := LogEntry{ClientID: "client-1", Timestamp: time.Now()}
	for i := 1; i <= 6; i++ {
		allowed := rL.Allow(entry)
		requestCount := rL.RequestCount(entry.ClientID)
		if allowed {
			fmt.Printf("request %d: allowed=%t, count=%d\n", i, allowed, requestCount)
		} else {
			fmt.Println("the request is not allowed because the request count ", requestCount, "is over the limit")
		}
	}
}
