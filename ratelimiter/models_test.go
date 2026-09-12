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

func TestRateLimiter_AllowWithWindow_ExpiresOldRequests(t *testing.T) {
	rL := NewRateLimiter()
	for i := 1; i <= 6; i++ {
		entry := LogEntry{ClientID: "client-1", Timestamp: time.Now()}
		allowed := rL.AllowCustom(entry, 100*time.Millisecond, 5)
		if i <= 5 && allowed != true {
			t.Errorf("request %d: got allowed=%t, want %t", i, allowed, true)
		}
		if i == 6 && allowed != false {
			t.Errorf("request %d: got allowed=%t, want %t", i, allowed, false)
		}
	}
	time.Sleep(150 * time.Millisecond)
	for i := 1; i <= 6; i++ {
		entry := LogEntry{ClientID: "client-1", Timestamp: time.Now()}
		allowed := rL.AllowCustom(entry, 100*time.Millisecond, 5)
		expected := i <= 5
		if allowed != expected {
			t.Errorf("request %d: got allowed=%t, want %t", i, allowed, expected)
		}
	}
}
