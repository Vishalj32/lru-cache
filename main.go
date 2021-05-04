package main

import (
	"fmt"
	lru "lru-cache/cache"
)

func main() {
	lru := lru.New(10)
	lru.Put(1, 10)
	lru.Put(2, 20)
	lru.Put(3, 30)

	fmt.Println(lru.Keys())
	fmt.Println(lru.Get(1))
	lru.Get(2)
	fmt.Println(lru.RecentlyUsed())
	lru.Remove(3)

	//lru.Purge()
	lru.Print()
}
