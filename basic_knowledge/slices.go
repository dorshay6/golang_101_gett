package main

import (
	"fmt"
)

func main() {
	aToF := []string{"A", "B", "C", "D", "E"}
	fmt.Println("aToF:", aToF)
	bToD := aToF[1:4]
	fmt.Println("bToD:", bToD)
	bToD[2] = "Dor"
	fmt.Println("aToF:", aToF)

	// This will overrite the data "E"
	bToD = append(bToD, "Fox")

	// Right now our cap = 4 but we will try to add a value
	// So we will create a new array for this slice and unlink from
	// aToF
	bToD = append(bToD, "Shay")
	bToD = append(bToD, "Zebra")
	fmt.Println("bToD:", bToD)

	//This cange will unlink it from aToF
	fmt.Println("aToF:", aToF)
}
