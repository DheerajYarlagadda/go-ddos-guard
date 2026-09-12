package main

import "time"

type LogEntry struct {
	ClientID  string
	Timestamp time.Time
}

type RateLimiter struct {
	requests map[string]int
}

func NewRateLimiter() *RateLimiter {
	r := make(map[string]int)
	return &RateLimiter{requests: r}
}

func (r *RateLimiter) RequestCount(clientID string) int {
	return r.requests[clientID]
}

func (r *RateLimiter) Allow(entry LogEntry) bool {
	requestCount := r.RequestCount(entry.ClientID)
	if requestCount >= 5 {
		return false
	}
	r.requests[entry.ClientID] = requestCount + 1
	return true
}
