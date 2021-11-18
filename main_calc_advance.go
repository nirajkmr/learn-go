package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	//reader := bufio.NewReader(os.Stdin)
	input1 := getInput("Enter value1:")
	input2 := getInput("Enter Value2:")

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Value 1 must be a number")
	// }
	// input2 = strings.TrimSpace(input2)
	// if err != nil {

	// 	fmt.Println(err)
	// 	panic("Value 2 must be integer")
	// }
	fmt.Println("Provide the Operation (+,-,*,/):")
	operator, _ := reader.ReadString('\n')
	var result float64
	switch operator {
	case "+":
		result = addValue(input1, input2)
	case "-":
		result = substractValue(input1, input2)
	case "*":
		result = multiplyValue(input1, input2)
	case "/":
		result = divideValue(input1, input2)
	default:
		fmt.Println("Invalid Operator Selection")
	}
}

func addValue(value1, value2 float64) float64 {

	return value1 + value2
}

func substractValue(value1, value2 float64) float64 {

	return value1 - value2
}
func multiplyValue(value1, value2 float64) float64 {

	return value1 * value2
}
func divideValue(value1, value2 float64) float64 {

	return value1 / value2
}

func getInput(prompt string) float64 {

	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	return value
}
