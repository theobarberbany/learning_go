// An exercise in pointers
package main

import "fmt"

func main() {
a, b := 1, 2 //initialise a and b
x, y := &b, &b // x and y both point to b

fmt.Printf("a has value %v, b has value %v\n", a, b)
fmt.Printf("Address of a is: %v, the address of b is: %v\n", &a, &b)
fmt.Printf("x holds value %v and has value %v \n", *x, x)
fmt.Printf("y has value %v and pointer %v \n", *y, &y)
fmt.Println("Setting x = 4")
*x=4
fmt.Printf("x has value %v and pointer %v \n", *x, &x)
fmt.Printf("y has value %v and pointer %v \n", *y, &y)
x = y
fmt.Println("Setting x = y")
fmt.Printf("x has value %v and pointer %v \n", *x, &x)
fmt.Printf("y has value %v and pointer %v \n", *y, &y)
var z *int
z = x
fmt.Printf("z has value %v and pointer %v \n", *z, &z)
b = a
fmt.Println("Set value of b to a")
fmt.Printf("x has value %v and pointer %v \n", *x, &x)
fmt.Printf("y has value %v and pointer %v \n", *y, &y)
fmt.Printf("z has value %v and pointer %v \n", *z, &z)
}

