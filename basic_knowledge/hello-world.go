// FILE: hello-world.go
package main

import (
	"fmt"
	"runtime"
)

// we import fmt and runtime, both of them are part of go
func main() {
	fmt.Println("You just run your first GO code on your", runtime.GOOS)
}
