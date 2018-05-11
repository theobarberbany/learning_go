package main

import (
	"log"
	"sync"
	"time"
)

// the function run inside goroutine. Recieves channel on ch, sleeps for t
// then sends t on channel it recieved
func doStuff(t time.Duration, ch <-chan chan time.Duration) {
	ac := <-ch
	time.Sleep(t)
	ac <- t
}

func main() {
	// create channel-over-channel type
	sendCh := make(chan chan time.Duration)

	//start up to 10 doStuff goroutines
	for i := 0; i < 10; i++ {
		go doStuff(time.Duration(i+1)*time.Second, sendCh)
	}

	// send channels to each doStuff goroutine. "ack" is achieved by sending sleep
	// time back.

	recvCh := make(chan time.Duration)
	for i := 0; i < 10; i++ {
		sendCh <- recvCh
	}

	// recieve on each channel previously sent. (recieve the ack)
	// WG willblock until all goroutines have recieved the ack and logged
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dur := <-recvCh
			log.Printf("Slept for %s", dur)
		}()
	}
	wg.Wait()
}
