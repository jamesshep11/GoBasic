package main

import "fmt"

func main() {
	var var1 string = "What are use"
	fmt.Printf("Value: %v\n", var1)
	fmt.Printf("Type: %T\n\n", var1)

	var var2 bool = false
	fmt.Printf("Value: %v\n", var2)
	fmt.Printf("Type: %T\n\n", var2)

	var var3 int8 = -128
	fmt.Printf("Value: %v\n", var3)
	fmt.Printf("Type: %T\n\n", var3)

	var var4 uint16 = 65535
	fmt.Printf("Value: %v\n", var4)
	fmt.Printf("Type: %T\n\n", var4)

	var var5 float32 = 255.63215645
	fmt.Printf("Value: %v\n", var5)
	fmt.Printf("Type: %T\n\n", var5)
}
