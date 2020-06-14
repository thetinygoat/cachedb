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

// LRU holds map and list which manage cache in an LRU fashion
type LRU struct {
	capacity uint64
	list     *List
	cache    map[string]*element
	size     uint64
}

// Init instantiates new LRU cache
func Init() *LRU {
	lru := &LRU{}
	lru.list = New()
	lru.cache = make(map[string]*element)
	lru.capacity = MaxMemory

	return lru
}

// Set inserts data into the cache
func (lru *LRU) Set(key string, value string) {
	if el, ok := lru.cache[key]; ok {
		lru.list.Remove(el)
		lru.list.Prepend(key, value)
		ref := lru.list.GetFirstRef()
		lru.cache[key] = ref
		oldMemUsage := len(el.value)
		newMemUsage := len(value)
		lru.size += uint64(newMemUsage - oldMemUsage)
	} else {
		if lru.size+uint64(len(value)) >= lru.capacity {
			data := lru.list.GetLast()
			if data == nil {
				return
			}
			lru.list.RemoveLast()
			delete(lru.cache, data[0])
			oldMemUsage := len(data[1])
			lru.size -= uint64(oldMemUsage)
		}
		lru.list.Prepend(key, value)
		ref := lru.list.GetFirstRef()
		lru.cache[key] = ref
		newMemUsage := len(value)
		lru.size += uint64(newMemUsage)
	}
}

// Get fetches a value related to a particular key from the cache
func (lru *LRU) Get(key string) string {
	el, ok := lru.cache[key]
	if !ok {
		return "nil"
	}
	k := el.key
	v := el.value
	lru.list.Remove(el)
	lru.list.Prepend(k, v)
	return v
}
