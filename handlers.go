package main

import (
	"strconv"
)

// Handler holds methods and configuration for request handling
type Handler struct {
	cache     Cache
	cacheType string
}

// CacheInit initializes the cache object with eviction policy specified in cachedb.conf
func CacheInit() *Handler {
	handler := &Handler{}
	if EvictionPolicy == "lru" {
		handler.cache = &LRU{list: New(), cache: make(map[string]*element), capacity: MaxMemory}
		handler.cacheType = "lru"
	} else if EvictionPolicy == "lazy" {
		handler.cache = &Lazy{cache: make(map[string]*cacheObject), capacity: MaxMemory}
		handler.cacheType = "lazy"
	}
	return handler
}

func (handler *Handler) handleGet(req []string) string {
	key := req[0]
	data, status := handler.cache.Get(key)
	if status == StatusUnexpected {
		panic(data)
	}

	return data
}

func (handler *Handler) handleSet(req []string) string {
	if len(req) < 2 {
		return MessageMalformed
	}
	key := req[0]
	value := req[1]
	var ttl int
	if len(req) == 3 {
		i, err := strconv.Atoi(req[2])
		if err != nil {
			return MessageMalformed
		}
		ttl = i
	} else {
		ttl = -1
	}
	data, status := handler.cache.Set(key, value, ttl)
	if status == StatusUnexpected {
		panic(MessageUnexpected)
	}
	return data
}
