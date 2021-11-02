package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func foo2() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo2", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
