package main

import (
	"fmt"
	"time"
)

func main() {
	ReadConfig()
	lru := &LRU{list: New(), cache: make(map[string]*element), capacity: MaxMemory}
	value, status := Set(lru, "sachin", "i am sachin", 2)
	fmt.Println(value, status)
	val, _ := Get(lru, "sachin")
	fmt.Println(val, lru.size)
	time.Sleep(4 * time.Second)
	val, _ = Get(lru, "sachin")
	fmt.Println(val, lru.size)
	fmt.Println(lru.cache, lru.capacity)

}
