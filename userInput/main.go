package main

import (
	"bufio"
	"fmt"
	"os"
)

// *****Reading user Input

// func main() {
// 	fmt.Println("hey! Whats your name")
// 	// var name string

// 	// fmt.Scan(&name)
// 	// fmt.Println("Hello, Mr.", name)

// 	reader := bufio.NewReader(os.Stdin)
// 	name, _ := reader.ReadString(naresh.txt)

// 	fmt.Println("Hello Mr.", name)

// }

//*************************Reading a file

func main() {
	file, err := os.Open("naresh.txt")

	if err != nil {
		fmt.Println("Error in opening", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	lineCount := 0
	for lineCount < 10 {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file:", err)
			}
			// this is to print the last line of the file
			fmt.Println(line)
			break
		}

		fmt.Println(line)
		lineCount++
	}

}
