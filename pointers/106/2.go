// An exercise in pointers
package main

import "fmt"

func main() {
a, b := 1, 2 //initialise a and b
x, y := &b, &b // x and y both point to b

fmt.Printf("a has value %v, b has value %v\n", a, b)
fmt.Printf("Address of a is: %v, the address of b is: %v\n", &a, &b)
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "x", *x, x, &x)
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "y", *y, y, &y)
fmt.Println("Setting x = 4")
*x=4
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "x", *x, x, &x)
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "y", *y, y, &y)
fmt.Println("Setting x = y = z")
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "x", *x, x, &x)
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "y", *y, y, &y)
var z *int
x=y
z=x
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "z", *z, z, &z)
b = a
fmt.Println("Set value of b to a")
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "x", *x, x, &x)
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "y", *y, y, &y)
fmt.Printf("pointer %q holds value %v and has value %v. It's memory address is %v \n", "z", *z, z, &z)
}

