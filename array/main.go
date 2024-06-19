package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Naresh"
	names[1] = "Kamlesh"
	names[2] = "Ramesh"
	names[3] = "Sitesh"
	names[4] = "Litesh"

	fmt.Println("Elements of names are:", names)
	fmt.Println("Length of the Array is:", len(names))

	var numbers = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers)

	var friends [5]string
	friends[0] = "Kitesh"
	fmt.Printf("Names of My friends are %q\n", friends[0])

}
