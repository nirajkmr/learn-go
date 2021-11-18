package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("math")
	i1, i2, i3 := 10, 15, 20
	intsum := i1 + i2 + i3
	fmt.Println("Add", intsum)
	f1, f2, f3 := 10.5, 11.34, 5.36
	sumFloat := f1 + f2 + f3
	fmt.Println("Float SUM", math.Round(sumFloat))
	circleRadius := 15.6
	circumfrence := circleRadius * 2 * math.Pi
	fmt.Printf("Print Circumference %2f", circumfrence)

}
