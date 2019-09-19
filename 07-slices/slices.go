package main

import (
	"fmt"
)

func main() {
	array := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(array)

	slice := array[1:4]
	fmt.Println(slice)

	slice = array[:2]
	fmt.Println(slice)

	slice = array[1:]
	fmt.Println(slice)
}

//len() e cap()
