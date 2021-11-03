package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	product := 1
	for product < 1000 {
		product *= 2
	}
	fmt.Println(product)

	count := 0
	for {
		count += 1
		break
	}
	fmt.Println(count)
}
