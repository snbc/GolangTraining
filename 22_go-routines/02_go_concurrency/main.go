package main

import "sync"

var total_foo int
var total_bar int
var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	println(total_bar)
	println(total_foo)

}

func foo() {
	for i := 0; i < 10; i++ {
		total_foo = total_foo + i
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		total_bar = total_bar + i
	}
	wg.Done()
}
