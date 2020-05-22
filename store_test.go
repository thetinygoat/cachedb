package main

import (
	"fmt"
	"testing"
	"time"
)

var testCacheObject = make(map[string]Cache)

func TestSet(t *testing.T) {
	var tests = []struct {
		key, value string
		ttl        int

		wantValue string
		wantTTL   time.Duration
	}{
		{"key1", "value1", 25, "value1", 25 * time.Second},
		{"key2", "value2", 135, "value2", 135 * time.Second},
		{"key3", "value3", 20, "value3", 20 * time.Second},
		{"key1", "value12", 31, "value12", 31 * time.Second},
		{"key1", "value13", 32, "value13", 32 * time.Second},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s -> %s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			testCacheObject = set(tt.key, tt.value, tt.ttl, testCacheObject)
			ans := testCacheObject[tt.key]
			if ans.Data != tt.wantValue || ans.TTL != tt.wantTTL {
				t.Errorf("error on key %s", tt.key)
			}
		})
	}
}
