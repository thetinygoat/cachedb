package main

import (
	"time"
)

// Cache : holds cache signature
type Cache struct {
	Data       string
	Identifier string
	CreatedAt  time.Time
	ExpireAt   time.Time
	TTL        time.Duration
}

func create(key string, value string, ttl int, globalCacheObject map[string]Cache) (map[string]Cache, error) {
	globalCacheObject[key] = Cache{
		Data:       value,
		Identifier: "id",
		CreatedAt:  time.Now(),
		ExpireAt:   time.Now().Add(time.Duration(ttl) * time.Second),
		TTL:        time.Duration(ttl) * time.Second,
	}
	return globalCacheObject, nil
}
