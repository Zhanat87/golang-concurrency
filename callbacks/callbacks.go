package main

import (
	"fmt"
	"sync"
	"strings"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	toUpperSync("hello callbacks", func(v string) {
		fmt.Printf("callback: %s\r\n", v)
		waitGroup.Done()
	})
	println("waiting async response")
	waitGroup.Wait()
}

func toUpperSync(word string, f func(string)) {
	go f(strings.ToUpper(word))
}
