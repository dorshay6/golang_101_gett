package main

import "fmt"

func main() {
	myArray := []string{
		"Gett",
		"Juno",
		"Lyft",
		"Uber",
	}

	for _, value := range myArray {
		fmt.Println("Loop A", value)
	loopB:
		for k := 0; k < 3; k++ {
			fmt.Println("Loop B", k)
			for f := 0; f < 2; f++ {
				if value != "Gett" && value != "Juno" {
					break loopB
				} else {
					fmt.Println("Loop C", f)
				}
			}
		}
	}
}
