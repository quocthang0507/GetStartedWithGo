package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("42 + 13 =", add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
