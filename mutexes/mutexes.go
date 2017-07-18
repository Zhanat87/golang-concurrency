// GOMAXPROCS=1 go run -race mutexes/mutexes.go
package main

import (
	"time"
	"sync"
)

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}
	for i := 0; i < 10; i++ {
		go func() {
			counter.Lock()
			counter.value++
			defer counter.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	counter.Lock()
	defer counter.Unlock()
	println(counter.value)
}
