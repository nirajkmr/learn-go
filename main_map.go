package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps")
	vehicle := make(map[string]string)
	fmt.Println(vehicle)
	vehicle["Car"] = "Audi"
	vehicle["Bike"] = "Royal Enfield"
	vehicle["Auto"] = "Bajaj"
	vehicle["Pickup"] = "Mahindra"
	fmt.Println(vehicle)
	auto := vehicle["Pickup"]
	fmt.Println(auto)
	delete(vehicle, "Bike")
	fmt.Println(vehicle)
	for k, v := range vehicle {
		fmt.Printf("%v: %v\n", k, v)
	}

}
