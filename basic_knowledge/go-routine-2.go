package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGorup sync.WaitGroup

func main() {
	waitGorup.Add(1)

	go printWorld()
	fmt.Print("Hello ")

	waitGorup.Wait()
}

func printWorld() {
	time.Sleep(3 * time.Second)
	fmt.Print("Dor")
	waitGorup.Done()
}
