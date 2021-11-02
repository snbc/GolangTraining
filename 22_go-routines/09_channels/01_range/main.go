package main

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for t := range c {
		print(t)
	}

}
