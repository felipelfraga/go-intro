package main

import (
	"fmt"
	"runtime"
)

func main() {

	if x := 1; x == 1 {
		fmt.Println("É igual a 1")
	} else {
		fmt.Println("Nunca será executado!")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Loop1: ", i)
	}

	i := 0
	for ; i < 5; i++ {
		fmt.Println("Loop2: ", i)
	}

	//for é o while do Go =)
	j := 0
	for j < 5 {
		fmt.Println("Loop3: ", j)
		j++
	}

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}


}
