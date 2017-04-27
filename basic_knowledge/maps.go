package main

import (
	"fmt"
)

func main() {
	taxiAppsInNY := map[string]int{
		"Gett": 1000,
		"Uber": 1200,
		"Juno": 400,
	}

	printMyMap(taxiAppsInNY)

	// Gett and Juno now are One!
	taxiAppsInNY["Gett"] = taxiAppsInNY["Gett"] + taxiAppsInNY["Juno"]
	delete(taxiAppsInNY, "Juno")

	fmt.Println("New world now !")
	printMyMap(taxiAppsInNY)
}

func printMyMap(myMap map[string]int) {
	for key, value := range myMap {
		fmt.Printf("App: %v has %v of rides in NYC every hour \n", key, value)
	}
}
