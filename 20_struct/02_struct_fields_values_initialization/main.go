package main

import "fmt"

type person struct {
	first string
	last  string
	int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.first, p1.last, p1.int)
	fmt.Println(p2.first, p2.last, p2.int)
}
