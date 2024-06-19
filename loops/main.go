package main

import "fmt"

func main() {
	// default simple for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Hello", i)
	// }

	// indefinite loop
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// condition loop (similar to while)
	// for i < 10 {
	// 	fmt.Println("Hello", i)
	// 	i++
	// }

	// on some data structure
	name := "Kamlesh"
	// number := [6]int{11, 12, 13, 14, 15, 16}
	for index, value := range name {
		fmt.Printf("number at index %d is %c\n", index, value)
	}

}
