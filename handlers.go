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
