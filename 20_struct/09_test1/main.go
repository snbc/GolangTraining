package main

import "fmt"

type Person struct {
	Name string
	age  int
}

type Student struct {
	Person
	StudentId int
}

func main() {
	stu := Student{}
	stu.StudentId = 4
	person := Person{"zhangsan", 34}
	stu.Person = person
	fmt.Printf("t%T", stu)
	fmt.Printf("%+v", stu)
}
