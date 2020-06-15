// Copyright (C) 2020 Sachin Saini

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"time"
)

// Cache specifies methods that other types could implement to be considered as a type of cache
type Cache interface {
	Set(string, string, int) (string, int)
	Get(string) (string, int)
}

type cacheObject struct {
	value   string
	ex      int
	addedAt time.Time
}

// Lazy implements cache with lazy cache eviction policy
type Lazy struct {
	cache    map[string]*cacheObject
	size     uint64
	capacity uint64
}

// Get returns value and status code for a particular key
func (lazy *Lazy) Get(key string) (string, int) {
	c, exists := lazy.cache[key]
	if !exists {
		return MessageNotFound, StatusNotFound
	}
	// indefinite expiration time
	if c.ex == -1 {
		return c.value, StatusOk
	}
	now := time.Now()
	ex := c.addedAt.Add(time.Duration(c.ex) * time.Second)
	// expired
	if now.Sub(ex) > 0 {
		lazy.size -= uint64(len(key) + len(c.value))
		delete(lazy.cache, key)
		return MessageExpired, StatusExpired
	}

	return c.value, StatusOk
}

// Set adds a key value pair to the cache
func (lazy *Lazy) Set(key string, value string, ex int) (string, int) {
	size := uint64(len(key) + len(value))
	if lazy.size+size > lazy.capacity {
		return "", StatusMemoryOverload
	}
	lazy.size += size
	lazy.cache[key] = &cacheObject{value: value, ex: ex, addedAt: time.Now()}
	return MessageOk, StatusOk
}

// LRU holds map and list which manage cache in an LRU fashion
type LRU struct {
	cache    map[string]*element
	size     uint64
	capacity uint64
	list     *List
}

// Set inserts data into the cache
func (lru *LRU) Set(key string, value string, ex int) (string, int) {
	if el, ok := lru.cache[key]; ok {
		lru.list.Remove(el)
		lru.list.Prepend(key, value, ex, time.Now())
		ref := lru.list.GetFirstRef()
		lru.cache[key] = ref
		oldMemUsage := len(el.value) + len(key)
		newMemUsage := len(value) + len(key)
		lru.size += uint64(newMemUsage - oldMemUsage)
		return MessageOk, StatusOk
	}
	if lru.size+uint64(len(value)+len(key)) >= lru.capacity {
		data := lru.list.GetLast()
		if data == nil {
			return MessageUnexpected, StatusUnexpected
		}
		lru.list.RemoveLast()
		delete(lru.cache, data[0])
		oldMemUsage := len(data[1])
		lru.size -= uint64(oldMemUsage)
	}
	lru.list.Prepend(key, value, ex, time.Now())
	ref := lru.list.GetFirstRef()
	lru.cache[key] = ref
	newMemUsage := len(value) + len(key)
	lru.size += uint64(newMemUsage)
	return MessageOk, StatusOk
}

// Get fetches a value related to a particular key from the cache
func (lru *LRU) Get(key string) (string, int) {
	el, exists := lru.cache[key]
	if !exists {
		return MessageNotFound, StatusNotFound
	}
	k := el.key
	v := el.value
	ex := el.ex
	addedAt := el.addedAt
	if ex == -1 {
		lru.list.Remove(el)
		lru.list.Prepend(k, v, ex, addedAt)
		return v, StatusOk
	}
	now := time.Now()
	exp := addedAt.Add(time.Duration(ex) * time.Second)
	if now.Sub(exp) > 0 {
		lru.list.Remove(el)
		delete(lru.cache, key)
		lru.size -= uint64(len(key) + len(v))
		return MessageExpired, StatusExpired
	}
	lru.list.Remove(el)
	lru.list.Prepend(k, v, ex, addedAt)
	return v, StatusOk
}

// Get forwrds request to specific handler
func Get(c Cache, key string) (string, int) {
	return c.Get(key)
}

// Set forwrds request to specific handler
func Set(c Cache, key string, value string, ex int) (string, int) {
	return c.Set(key, value, ex)
}
