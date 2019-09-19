package main

import (
	"fmt"
	"time"
)

func vai(n int) {
	fmt.Println(n)
}

func main() {
	go vai(10)
	vai(11)
	time.Sleep(100 * time.Millisecond)
}
