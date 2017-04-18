// FILE: variables.go
package main

import "fmt"

func main() {
	// Shorthand declartion
	userStatusText := "User never used our app"
	fmt.Println("[main] status: ", userStatusText)
	// This will send the value of userStatusText to function, original wont be changed
	changeUserStatusAfterUseValue(userStatusText)
	fmt.Println("[main] status: ", userStatusText)
	// Now we will send refernce of userStatusText insted of value
	changeUserStatusAfterUseRefernce(&userStatusText)
	fmt.Println("[main] status: ", userStatusText)
}

func changeUserStatusAfterUseValue(userStatusText string) string {
	userStatusText = "User used the app twice"
	fmt.Println("[changeUserStatusAfterUseValue] status: ", userStatusText)
	return userStatusText
}

func changeUserStatusAfterUseRefernce(userStatusText *string) string {
	*userStatusText = "User used the app twice"
	fmt.Println("[changeUserStatusAfterUseRefernce] status: ", *userStatusText)
	return *userStatusText
}
