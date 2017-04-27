package main

import (
	"fmt"
)

func main() {
	type TaxiDriver struct {
		name      string
		rank      int
		carNumber string
	}

	newDriver := TaxiDriver{
		name:      "Dor Shay",
		rank:      4,
		carNumber: "308141482",
	}
	// you can also use `newDriver := new(TaxiDriver)` and get the pointer
	// Or use `var newDriver TaxiDriver`

	fmt.Println("Your driver name is:", newDriver.name)
}
