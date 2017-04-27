package main

import (
	"fmt"
	"time"
)

func main() {
	go printWorld()
	fmt.Print("Hello ")

	// We need main to no exit on us, soo just like go rutine
	// we will use sleep to keep it alive
	time.Sleep(2 * time.Second)
}

func printWorld() {
	time.Sleep(3 * time.Second)
	fmt.Print(" Dor")
}
