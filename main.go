package main

import (
	"fmt"
	"time"
)

func main() {
	rL := NewRateLimiter()
	for i := 1; i <= 6; i++ {
		entry := LogEntry{ClientID: "client-1", Timestamp: time.Now()}
		allowed := rL.Allow(entry)
		requestCount := rL.RequestCount(entry.ClientID)
		if allowed {
			fmt.Printf("request %d: time %v allowed=%t, count=%d\n", i, entry.Timestamp, allowed, requestCount)
		} else {
			fmt.Println("the request is not allowed because the request count ", requestCount, "is over the limit")
		}
	}
}
