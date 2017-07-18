package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	goroutines := 5
	waitGroup.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			fmt.Printf("#%d goroutine: hello world!\r\n", i)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}
