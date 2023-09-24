package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, error := strconv.Atoi("sdfgf")

	fmt.Printf("Result: %v\n", result)

	if error != nil {
		fmt.Printf("Error: %v", error)
	}
}
