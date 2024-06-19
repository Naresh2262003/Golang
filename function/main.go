package main

import "fmt"

func simpleFunction(a int, b int) int {
	result := a + b
	return result
}

func square(a int) (result int) {
	result = a * a
	return
}

func main() {
	a := 5
	b := 10

	fmt.Println("We are learning functions in Golang.")
	fmt.Printf("Sum of %d and %d is %d", a, b, simpleFunction(a, b))

	fmt.Printf("The square of the %d is %d", a, square(a))
}
