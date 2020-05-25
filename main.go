package main

import "fmt"

func main() {
	list := List{Name: "mylist"}
	list.rpush("1")
	list.rpush("2")
	list.rpush("3")
	list.rpush("4")
	list.rpush("5")
	list.rpush("6")
	fmt.Println(list.lpop())
	fmt.Println(list.lpop())
	fmt.Println(list.lpop())
	fmt.Println(list.lpop())
	fmt.Println(list.lpop())
	fmt.Println(list.lpop())
	list.lrange(0, 999)
}
