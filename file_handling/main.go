package main

import (
	"fmt"
	"os"
)

// 1) create a file and insert data in that
// func main() {
// 	file, err := os.Create("example.txt")

// 	if err != nil {
// 		fmt.Println("Error while creating a file", err)
// 		return
// 	}

// 	defer file.Close()

// 	content := "Hello this is Naresh \nlooking for something fun around "
// 	byte, errs := io.WriteString(file, content+"\n")

// 	if errs != nil {
// 		fmt.Println("error while writing in the file", errs)
// 		return
// 	}

// 	fmt.Println("successfully created file", byte)

// }

// 2) Readong a file using buffer  (Useful when to read big file as reads the data chunk by chunk)
// func main() {
// 	file, err := os.Open("example.txt")

// 	if err != nil {
// 		fmt.Println("Error while creating a file", err)
// 		return
// 	}
// 	defer file.Close()

// 	// creating a buffer to read the file content
// 	buffer := make([]byte, 1024)

// 	// Read the file content into the file
// 	for {
// 		n, err := file.Read(buffer)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			fmt.Println("Error while reading the file", err)
// 			return
// 		}

// 		fmt.Println(string(buffer[:n]))

// 	}

// 	fmt.Println("Done")
// }

// 3) Readong a file using function
func main() {
	content, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("error while reading the file", err)
		return
	}

	fmt.Println(string(content))
}
