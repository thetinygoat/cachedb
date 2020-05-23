package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	testCacheRef := GlobalCacheRef{}
	testCacheRef.intializeGlobalCache()
	var tests = []struct {
		key, value string
		ttl        int

		wantValue string
		wantTTL   time.Duration
	}{
		{"key1", "value1", -1, "value1", -1 * time.Second},
		{"key2", "value2", 135, "value2", 135 * time.Second},
		{"key3", "value3", 20, "value3", 20 * time.Second},
		{"key1", "value12", 31, "value12", 31 * time.Second},
		{"key1", "value13", 32, "value13", 32 * time.Second},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s->%s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			testCacheRef.set(tt.key, tt.value, tt.ttl)
			ans := testCacheRef.Cache[tt.key]
			if ans.Data != tt.wantValue || ans.TTL != tt.wantTTL {
				t.Errorf("error on key %s", tt.key)
			}
		})
	}
}

func seed(testCacheRef GlobalCacheRef) {
	var data = []struct {
		key, value string
		ttl        int
	}{
		{"key1", "value1", -1},
		{"key2", "value2", 20},
		{"key3", "value3", 20},
	}
	for _, d := range data {
		testCacheRef.set(d.key, d.value, d.ttl)
	}
}
func TestGet(t *testing.T) {
	testCacheRef := GlobalCacheRef{}
	testCacheRef.intializeGlobalCache()
	var tests = []struct {
		key   string
		sleep int

		want string
	}{
		{"key1", 30, "value1"},
		{"key2", 25, KeyExpired},
		{"key3", 0, "value3"},
		{"key4", 0, KeyDoesNotExist},
	}
	var data = []struct {
		key, value string
		ttl        int
	}{
		{"key1", "value1", -1},
		{"key2", "value2", 20},
		{"key3", "value3", 20},
		{"key5", "value5", 90},
	}
	for i := range tests {
		testname := fmt.Sprintf("%s", tests[i].key)
		t.Run(testname, func(t *testing.T) {
			testCacheRef.set(data[i].key, data[i].value, data[i].ttl)
			time.Sleep(time.Duration(tests[i].sleep) * time.Second)
			ans, err := testCacheRef.get(tests[i].key)
			if err != nil {
				if err.Error() != tests[i].want {
					t.Errorf("error on key %s", tests[i].key)
				}
			} else if ans != tests[i].want {
				t.Errorf("error on key %s", tests[i].key)
			}
		})
	}
}
