// go run channels/channels.go
package main

import (
	"fmt"
	"time"
	"strconv"
)

func main1() {
	channel := make(chan string)
	go func(){
		channel <- "message to channel"
	}()
	message := <-channel
	fmt.Printf("%s\r\n", message)
}

func main2() {
	channel := make(chan string, 1)
	go func(){
		channel <- "message to channel 1"
		channel <- "message to channel 2"
	}()
	message := <-channel
	fmt.Printf("%s\r\n", message)
}

func main3() {
	channel := make(chan string, 1)
	go func(ch chan<- string){
		ch <- "message to channel 1"
		time.Sleep(1 * time.Second)
		println("finishing goroutine")
	}(channel)
	// note: channel not wait until goroutine was finished,
	// channel will get value and it is all
	message := <-channel
	fmt.Printf("%s\r\n", message)
}

func recievingCh(ch <-chan string) {
	msg := <-ch
	println(msg)
	// note: error was here, because it was only receiving channel
	//ch <- "hello"
}

func main4() {
	helloCh := make(chan string, 1)
	goodByeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiver(helloCh, goodByeCh, quitCh)
	go sendString(helloCh, "hello!")
	time.Sleep(1 * time.Second)
	go sendString(goodByeCh, "good bye!")
	// note: this quitCh alson known as forever channel, will wait until
	// not received value
	<-quitCh
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodByeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodByeCh:
			println(msg)
		case <-time.After(2 * time.Second):
			println("Nothing received in 2 seconds. exit")
			quitCh <- true
			break
		}
	}
}

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		close(ch)
	}(ch)

	for v := range ch {
		println("ch range: " + strconv.Itoa(v))
	}
}
