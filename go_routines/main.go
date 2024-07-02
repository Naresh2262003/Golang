package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hello End ")
}

func sayHi() {
	fmt.Println("Hi World")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hi End")
}

func main() {
	fmt.Println("Learning go routines")
	go sayHello()
	go sayHi()

	time.Sleep(1010 * time.Millisecond)
}
