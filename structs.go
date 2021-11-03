package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{2, 1})

	v := Vertex{2, 1}
	v.X = 4
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2} // has type Vertex
		v2 = Vertex{X: 1} // Y:0 is implicit
		v3 = Vertex{}     // X:0 and Y:0
		p1 = &Vertex{1, 2}
	)
	fmt.Println(v1, p1, v2, v3)
}
