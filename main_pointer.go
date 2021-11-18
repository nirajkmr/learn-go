package main

import (
	"fmt"
)

func main() {

	var ptr *int
	ptr2 := 42
	ptr = &ptr2

	fmt.Println("Pointer:", *ptr)
	fmt.Println("Pointer:", ptr2)
}
