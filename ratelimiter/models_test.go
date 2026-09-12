package ratelimiter

import (
	"testing"
	"time"
)

func TestRateLimiter_Allow_BlocksAfterLimit(t *testing.T) {
	rL := NewRateLimiter()
	for i := 1; i <= 6; i++ {
		entry := LogEntry{ClientID: "client-1", Timestamp: time.Now()}
		allowed := rL.Allow(entry)
		expected := i <= 5
		if allowed != expected {
			t.Errorf("request %d: got allowed=%t, want %t", i, allowed, expected)
		}
	}
}
