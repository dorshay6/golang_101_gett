package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var m map[string]json.RawMessage
	b := []byte(`{"dor": 1, "shay": 2}`)
	err := json.Unmarshal(b, &m)

	if err != nil {
		fmt.Println("Err", err)
	}
	fmt.Println("Result ", m)

	b = []byte(`{dor": 1, "shay": 2}`)
	err = json.Unmarshal(b, &m)

	if err != nil {
		fmt.Println("Err", err)
	}
}
