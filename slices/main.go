package main

import "fmt"

func main() {
	// This is array- fixed size
	// numbers := [5]int{1, 2, 3, 4, 5}

	// This is slice- variable size
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 13)
	fmt.Printf("Data Type of Numbers is %d\n", numbers)
	fmt.Printf("The length of slice is %d\n", len(numbers))

	// creating empty slice
	// marks := []int{}

	// creating empty slice
	marks := make([]int, 0, 5)
	marks = append(marks, 0, 1, 2, 3, 4, 5, 6, 4, 7, 8, 9, 4, 0, 02, 32, 25, 65, 25, 25, 45, 74)

	fmt.Println("Elements of marks are :", marks)
	fmt.Println("Length of marks is:", len(marks))
	fmt.Println("Capacity of marks is:", cap(marks))

}
