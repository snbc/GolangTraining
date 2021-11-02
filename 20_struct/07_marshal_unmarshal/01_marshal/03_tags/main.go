package main

import "fmt"

type person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

type person2 struct {
	First string
	Last  string `json:"last"`
	Age   int    `json:"wisdom score"`
}

func main() {
	//p1 := person2{"James", "Bond", 20}
	//bs, _ := json.Marshal(p1)
	//fmt.Println(string(bs))
	p1 := person2{"James", "Bond", 20}
	fmt.Println(p1)
	fmt.Println(p1)
}
