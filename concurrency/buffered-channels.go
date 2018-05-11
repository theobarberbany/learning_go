package main

import "fmt"
//channel only blocks when buffer is full, second argument to make is buffer size
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
