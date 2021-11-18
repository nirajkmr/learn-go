package main

import "fmt"

func main() {

	theAnswer := 42
	var result string
	if theAnswer > 0 {
		result = "less than zero"
	} else if theAnswer == 0 {
		result = "equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Print(result)

}
