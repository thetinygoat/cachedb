package main

import "fmt"

func main() {
	var globalCacheObject = make(map[string]Cache)
	globalCacheObject, _ = create("key1", "value1", 15, globalCacheObject)
	globalCacheObject, _ = create("key2", "value2", 20, globalCacheObject)
	fmt.Println(globalCacheObject)
}
