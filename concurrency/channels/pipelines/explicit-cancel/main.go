package main

import "fmt"
import "sync"

// Change to buffered output channel
func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// done recieves struct{} as the value does not matter, but the recieve does.
	var wg sync.WaitGroup
	out := make(chan int)

	// Start output goroutine for each input channel in cs. output
	// copies values from c to out until c is closed, or it recieves a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select { // select is blocking, hence a case being selected proceeds?
			case out <- n: // on the recieve action
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("vim-go")
	// done channel shared by whole pipeline
	// close channel when pipeline exits
	// that 'close' signals all goroutines to also exit
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)
	// Distribute sq work across two goroutines, both reading from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume first value from output

	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 || 9

	// done closed  be deferred call
}
