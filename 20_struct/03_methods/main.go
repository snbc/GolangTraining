package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	fmt.Println(p)
	return p.first + p.last
}
func (p *person) myage() int {
	fmt.Println(*p)
	return (*p).age
}

func main() {
	p1 := person{"James", "Bond", 20}
	//p2 := person{"Miss", "Moneypenny", 18}
	//fmt.Println(p1.fullName())
	//fmt.Println(p2.fullName())
	//fmt.Print(p1.fullName())
	//fmt.Print(p2.myage())
	p1.fullName()
	p1.myage()

}
