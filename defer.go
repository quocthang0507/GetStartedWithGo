package main

import "fmt"

func stack() {
	fmt.Println("counting")
	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are
	// executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func print() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("hello")
}

func main() {
	print()

	stack()
}
