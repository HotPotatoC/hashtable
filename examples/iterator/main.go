package main

import (
	"fmt"

	"github.com/HotPotatoC/hashtable"
)

func main() {
	ht := hashtable.New()

	for i := 0; i < 10; i++ {
		ht.Set(fmt.Sprintf("k%d", i+1), i+1)
	}

	for entry := range ht.Iter() {
		fmt.Printf("key: %s | value: %v\n", entry.Key, entry.Value)
	}
}
