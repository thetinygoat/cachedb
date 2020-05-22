package main

import (
	"fmt"
)

func main() {
	var globalCacheObject = make(map[string]Cache) // global cache ref
	globalCacheObject = set("key1", "value1", 15, globalCacheObject)
	globalCacheObject = set("key2", "value2", 20, globalCacheObject)
	fmt.Println(globalCacheObject)
	value, err := get("key2", globalCacheObject)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
	fmt.Println(globalCacheObject)
}
