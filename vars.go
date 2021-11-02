package main

import "fmt"

var c, python, java bool
var s string
var i, j int = 1, 2

func main() {
	var n int
	fmt.Println(n, c, python, java)

	c, python, java, s = true, false, true, "no!"
	fmt.Println(i, j, c, python, java, s)

	k := 3 // implicit type
	fmt.Println(k)
}
