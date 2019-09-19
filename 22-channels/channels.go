package main

import (
	"fmt"
	"time"
)

func vai(i int, c chan int) {
	time.Sleep(500 * time.Millisecond)
	c <- i * 2
}

func main() {

	c := make(chan int)
	go vai(10, c)
	go vai(20, c)
	i, j := <-c, <-c
	fmt.Println(i, j)

}
