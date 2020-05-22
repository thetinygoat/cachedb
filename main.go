package main

import (
	"fmt"
)

func main() {
	var globalCacheObject = make(map[string]Cache) // global cache ref
	globalCacheObject = set("key1", "value1", 15, globalCacheObject)
	fmt.Println(globalCacheObject)
	flush(globalCacheObject)
	fmt.Println(globalCacheObject)

}
