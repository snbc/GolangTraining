package main

import "fmt"

func main() {
	//test fsprintf
	//fmt.Println(printurl())
	printPoint()

}

type point struct {
	x, y int
}

func printurl() string {
	var socketcode = "hello,world"
	var name = "hsj"
	var url = fmt.Sprintf("socketcode=%s&name=%s", socketcode, name)
	return url
}

func printPoint() {
	p := point{1, 3}
	fmt.Printf("%v\n", p)

}
