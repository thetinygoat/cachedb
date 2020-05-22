package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// errors
const (
	KeyDoesNotExist = "KEY_DOES_NOT_EXIST_ERROR"
	KeyExpired      = "KEY_EXPIRED_ERROR"
)

// Cache : holds cache signature
type Cache struct {
	Data       string
	Identifier string
	CreatedAt  time.Time
	ExpireAt   time.Time
	TTL        time.Duration
}

// insert data in the store
func set(key string, value string, ttl int, globalCacheObject map[string]Cache) map[string]Cache {
	globalCacheObject[key] = Cache{
		Data:       value,
		Identifier: uuid.New().String(),
		CreatedAt:  time.Now(),
		ExpireAt:   time.Now().Add(time.Duration(ttl) * time.Second),
		TTL:        time.Duration(ttl) * time.Second,
	}
	return globalCacheObject
}

func get(key string, globalCacheObject map[string]Cache) (string, error) {
	if cacheObject, ok := globalCacheObject[key]; ok {
		if time.Now().Sub(cacheObject.ExpireAt) > 0 {
			delete(globalCacheObject, key)
			return KeyExpired, errors.New(KeyExpired)
		}
		return cacheObject.Data, nil
	}
	return KeyDoesNotExist, errors.New(KeyDoesNotExist)

}

func del(key string, globalCacheObject map[string]Cache) error {
	if _, ok := globalCacheObject[key]; ok {
		delete(globalCacheObject, key)
		return nil
	}
	return errors.New(KeyDoesNotExist)
}
