package main

import (
	"fmt"
)

func main() {

	doSomthing()
	sum := addValues(10, 5)
	fmt.Println("Sum of the value", sum)
	allSum, length := addAllValue(10, 5, 20)
	fmt.Println(allSum)
	fmt.Print(length)
}

func doSomthing() {

	fmt.Print("do somthing")

}

func addValues(Value1 int, Value2 int) int {

	return Value1 + Value2
}

func addAllValue(Values ...int) (int, int) {
	total := 0
	for _, v := range Values {
		total += v
	}
	return total, len(Values)
}
