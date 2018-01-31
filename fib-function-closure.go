package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fn,fn1,fn0 := 0,1,0
	return func () int {
		fn,fn1,fn0 = fn0,fn0+fn1,fn1
		return fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

