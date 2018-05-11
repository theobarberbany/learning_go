package main

import (
	"log"
	"sync"
	"time"
)

// struct to be passed over a channel to goroutine
type process struct {
	dur time.Duration
	ch  chan time.Duration
}

// goroutine func, recieves process struct 'p' on ch,
// sleeps for p.dur then send p.dur on p.ch
func doStuff(ch <-chan process) {
	proc := <-ch
	time.Sleep(proc.dur)
	proc.ch <- proc.dur
}

func main() {
	// start goroutines
	sendCh := make(chan process)
	for i := 0; i < 10; i++ {
		go doStuff(sendCh)
	}

	// store array of each struct sent to a goroutine
	processes := make([]process, 10)
	for i := 0; i < 10; i++ {
		dur := time.Duration(i+1) * time.Second
		proc := process{dur: dur, ch: make(chan time.Duration)}
		processes[i] = proc
		sendCh <- proc
	}

	// recieve on each struct's ack channel
	// wg blocks until all goroutines have acked and logged
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ch <-chan time.Duration) {
			defer wg.Done()
			dur := <-ch
			log.Printf("slept for %s", dur)
		}(processes[i].ch)
	}
	wg.Wait()
}
