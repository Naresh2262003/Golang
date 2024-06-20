package main

import "fmt"

func modifyValueByReference(ref *int) {
	*ref = *ref + 20
}

func main() {
	fmt.Println()

	num := 25

	pointer := &num

	fmt.Println("Address : ", pointer)
	fmt.Printf("Value at Address %d is %d\n", pointer, *pointer)

	// var ptr *int
	// if ptr == nil {
	// 	fmt.Println("ptr is nil")
	// }

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains :", value)
}
