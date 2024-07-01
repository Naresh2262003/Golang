package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("Live in like th")

	person := Person{Name: "Naresh", Age: 34, IsAdult: true}
	fmt.Println("person Data is ", person)

	// convert person into json encoding (Marshalling)
	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}
	fmt.Println("Person Data is ", string(jsonData))

	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Person data after unmarshalling", personData)
}
