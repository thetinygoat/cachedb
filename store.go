package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// constants
const (
	KeyDoesNotExist = "KEY_DOES_NOT_EXIST_ERROR"
	KeyExpired      = "KEY_EXPIRED_ERROR"
	OK              = "OK"
)

// CacheObject : holds cache signature
type CacheObject struct {
	Data       string
	Identifier string
	CreatedAt  time.Time
	ExpireAt   time.Time
	TTL        time.Duration
}

// GlobalCacheRef : maintains global cache state
type GlobalCacheRef struct {
	Cache map[string]CacheObject
}

func (c *GlobalCacheRef) intializeGlobalCache() {
	c.Cache = make(map[string]CacheObject)
}

// insert data in the store
func (c *GlobalCacheRef) set(key string, value string, ttl int) string {
	c.Cache[key] = CacheObject{
		Data:       value,
		Identifier: uuid.New().String(),
		CreatedAt:  time.Now(),
		ExpireAt:   time.Now().Add(time.Duration(ttl) * time.Second),
		TTL:        time.Duration(ttl) * time.Second,
	}
	return OK
}

// retrieve data from the store
func (c *GlobalCacheRef) get(key string) (string, error) {
	if cache, ok := c.Cache[key]; ok {
		// check for infinite expire time
		if cache.TTL < 0 {
			return cache.Data, nil
		}
		// check for key expiration
		if time.Now().Sub(cache.ExpireAt) > 0 {
			delete(c.Cache, key)
			return "", errors.New(KeyExpired)
		}
		return cache.Data, nil
	}
	return "", errors.New(KeyDoesNotExist)

}

// delete a key from the store
func (c *GlobalCacheRef) del(key string) (string, error) {
	if _, ok := c.Cache[key]; ok {
		delete(c.Cache, key)
		return OK, nil
	}
	return "", errors.New(KeyDoesNotExist)
}

// clear the global cache store
func (c *GlobalCacheRef) flush() string {
	for key := range c.Cache {
		delete(c.Cache, key)
	}
	return OK
}
