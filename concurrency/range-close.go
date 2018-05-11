package main

import (
	"fmt"
)

// syntax: v, ok := <-ch
// returns ok = false when ch is closed.


//close used to inform the range loop to terminate
//sender close, never receiver 
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

