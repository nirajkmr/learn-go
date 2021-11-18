package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You Entered: ", input)
	fmt.Println("Enter a number")
	numInput, asx := reader.ReadString('\n')
	afloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	fmt.Print(asx)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("You Entered: ", afloat)
	}

}
