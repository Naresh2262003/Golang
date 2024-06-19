package main

import "fmt"

func main() {
	age := 2

	if age >= 60 {
		fmt.Println("You are a senior citizen")
	} else if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}
}
