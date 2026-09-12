package ratelimiter

import "time"

type LogEntry struct {
	ClientID  string
	Timestamp time.Time
}

type RateLimiter struct {
	requests map[string][]time.Time
}

func NewRateLimiter() *RateLimiter {
	r := make(map[string][]time.Time)
	return &RateLimiter{requests: r}
}

func (r *RateLimiter) ClientRequests(clientID string) []time.Time {
	return r.requests[clientID]
}

func (r *RateLimiter) RecentRequests(clientID string, window time.Duration) []time.Time {
	recentRequests := make([]time.Time, 0)
	for _, request := range r.ClientRequests(clientID) {
		requestDuration := time.Since(request)
		if window > requestDuration {
			recentRequests = append(recentRequests, request)
		}
	}
	return recentRequests
}

func (r *RateLimiter) RequestCount(clientID string) int {
	return len(r.ClientRequests(clientID))
}

func (r *RateLimiter) RecentRequestCount(clientID string, window time.Duration) int {
	return len(r.RecentRequests(clientID, window))
}

func (r *RateLimiter) Allow(entry LogEntry) bool {
	return r.AllowCustom(entry, 2*time.Second, 5)
}

func (r *RateLimiter) AllowCustom(entry LogEntry, window time.Duration, maxRequests int) bool {
	recentRequestCount := r.RecentRequestCount(entry.ClientID, window)
	if recentRequestCount >= maxRequests {
		return false
	}
	r.requests[entry.ClientID] = append(r.requests[entry.ClientID], entry.Timestamp)
	return true
}
