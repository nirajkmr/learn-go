package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr = []string{"Red", "Blue", "Green"}
	fmt.Println(arr)
	arr = append(arr, "Yellow")
	fmt.Println(arr)
	number := make([]int, 5)
	number[0] = 134
	number[1] = 135
	number[2] = 136
	number[3] = 138
	number[4] = 137
	fmt.Print(number)
	sort.Ints(number)
	fmt.Print(number)
}
