package main

import (
	"fmt"
	"sync/atomic"

	"github.com/HotPotatoC/hashtable"
)

func main() {
	ht := hashtable.New()
	var num int64

	ht.Set("counter", &num)

	data, _ := ht.Get("counter")
	counter := data.(*int64)

	atomic.AddInt64(counter, 100)
	atomic.AddInt64(counter, 100)
	atomic.AddInt64(counter, 100)

	value := atomic.LoadInt64(counter)
	fmt.Println(value)
}
