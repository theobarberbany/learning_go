package main

import (
	"log"
	"sync"
	"time"
)

// doStuff knows nothing about a return channel
// code can be changed later to do whatever with ackFn
func doStuff(dur time.Duration, ch <-chan func(time.Duration)) {
	ackFn := <-ch
	time.Sleep(dur)
	ackFn(dur)
}

func main() {
	// start up doStuff goroutines
	sendCh := make(chan func(time.Duration))
	for i := 0; i < 10; i++ {
		dur := time.Duration(i+1) * time.Second
		go doStuff(dur, sendCh)
	}

	// Make channels that will be closed over, create  functions
	// that close over each channel, then send them to doStuff goroutines
	recvChs := make([]chan time.Duration, 10)
	for i := 0; i < 10; i++ {
		recvCh := make(chan time.Duration)
		recvChs[i] = recvCh
		fn := func(dur time.Duration) {
			recvCh <- dur
		}
		sendCh <- fn
	}

	// recieve on closed-over functions
	// wg blocks until all goroutines have recieved the ack and logged
	var wg sync.WaitGroup
	for _, recvCh := range recvChs {
		wg.Add(1)
		go func(recvCh <-chan time.Duration) {
			defer wg.Done()
			dur := <-recvCh
			log.Printf("slept for %s", dur)
		}(recvCh)
	}
	wg.Wait()
}
