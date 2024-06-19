package main

// %s - string
// %d - Integer
// %T - to get the type of the variable
// %.2f - to get the value till 2 points (float)

import "fmt"

func main() {
	var chem string
	fmt.Println("Mongo is love")
	chem = "Naresh"
	fmt.Println("I am a man ", chem)

	age := 25
	name := "Naresh"
	height := 5.232

	fmt.Printf("Type of age is %T\n", height)

	fmt.Printf("Hi i am %s and my age is %d and My height in meters is %.2f \n", name, age, height)
}
