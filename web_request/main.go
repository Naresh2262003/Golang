package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning with service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting Get response ", err)
		return
	}
	defer res.Body.Close()
	// fmt.Printf("Type of response %T\n", res)

	// Read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}

	fmt.Println("response:", string(data))

}
