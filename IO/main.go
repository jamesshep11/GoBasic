package main

import "fmt"

func main() {
	ting := getInput()
	writeOutput(ting)
}

func getInput() string {
	var in string

	fmt.Println("Gimme summin:")
	fmt.Scanln(&in)

	return in
}

func writeOutput(output string) {
	fmt.Printf("Interesting. Why %v?", output)
}
