package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Value1: ")
	input, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	fmt.Println("Enter the Value2:")
	numInput, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {

		fmt.Println(err)
		panic("Value 2 must be integer")
	}
	sumInput := float1 + float2
	sumInput = math.Round(sumInput*100) / 100
	fmt.Printf("The Sum of %v and %v is %v\n\n", float1, float2, sumInput)

}
