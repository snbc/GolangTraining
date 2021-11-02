package main

import "fmt"

type square struct {
	side float64
}

type square2 struct {
	side float64
}

func (z square2) area2() float64 {
	return z.side * z.side
}

func (z square) area() float64 {
	return z.side * z.side
}

func main() {
	//s := square{10}
	//fmt.Println("Area: ", s.area())
	s2 := square2{10}
	fmt.Println("Area", s2.area2())
}
