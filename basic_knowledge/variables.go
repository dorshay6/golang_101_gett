// FILE: variables.go
package main

import (
	"fmt"
	"reflect"
)

var (
	// Multiplate varibels
	driverName, userName string
	orderedClassNameID   int
	// Set type scroding to initaltion
	locale = "IL"
)

// we import fmt and runtime, both of them are part of go
func main() {
	// Shorthand declartion
	orderedClassName := "Gett Express"
	fmt.Println("driverName : ", driverName, " and typed as: ", reflect.TypeOf(driverName))
	fmt.Println("userName : ", userName, " and typed as: ", reflect.TypeOf(userName))
	fmt.Println("orderedClassNameID : ", orderedClassNameID, " and typed as: ", reflect.TypeOf(orderedClassNameID))
	fmt.Println("locale : ", locale, " and typed as: ", reflect.TypeOf(locale))
	fmt.Println("orderedClassName : ", orderedClassName, " and typed as: ", reflect.TypeOf(orderedClassName))
}
