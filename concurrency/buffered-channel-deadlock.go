package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 4
	ch <- 2
	ch <- 0
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
