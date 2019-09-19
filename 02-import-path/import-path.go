package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	msg := stringutil.Reverse("hello, world")
	fmt.Println(msg)
}
