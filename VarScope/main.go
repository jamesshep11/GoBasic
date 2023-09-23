package main

import "fmt"

const stay string = "This stays the same"

var global string = "This just for us"
var Public string = "This is for everyone"

func main() {
	// Classic declaration
	var word string
	word = "This is a string"
	fmt.Printf("%v\n", word)

	// Declare and assign
	var thing string = "This is a thing"
	fmt.Printf("%v\n", thing)

	// Implicit declaration
	var stuff = "This is some stuff"
	fmt.Printf("%v\n", stuff)

	// No var style (not for public)
	private := "This is mine"
	fmt.Printf("%v\n", private)

	global = "This just for us"
	fmt.Printf("%v\n", global)

	Public = "This is for everyone"
	fmt.Printf("%v\n", Public)

	fmt.Printf("%v\n", stay)
}
