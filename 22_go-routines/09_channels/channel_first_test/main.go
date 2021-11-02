package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("groutinue begion")
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
			fmt.Println("groutinue end")
		}

	}()

	for data := range ch {
		fmt.Println(data)
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			break
		}
	}

}
