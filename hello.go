package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}
