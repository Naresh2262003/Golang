package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error", err)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	// fmt.Println("Data", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("data", todo)

	fmt.Println(todo.Completed)
}

func postRequest() {
	todo := Todo{
		UserId:    2,
		Title:     "Naresh",
		Completed: true,
	}

	// convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert the string into Reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("Response", string(data))
}

func updateRequest() {
	todo := Todo{
		UserId:    215,
		Title:     "Lily",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while ", err)
		return
	}

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// create PUT Request
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Updated", string(data))
	fmt.Println("Status", res.StatusCode)
}

func deleteRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// create DELETE Req

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error creatin Delete Request", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creatin Delete Request", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Status ", res.Status)
}

func main() {
	fmt.Println("Learning CRUD")

	getRequest()
	postRequest()
	updateRequest()
	deleteRequest()
}
