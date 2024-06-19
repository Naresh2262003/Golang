package main

import (
	"fmt"
	color "mylearning/colorblindness"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello World!!")
	myutil.Position()
	color.Mycolor()

	var name string = "Naresh"
	fmt.Printf("My name is %s\n", name)

	// var Prince = 10

	var version = "Kamlesh"

	var money int = 2500
	var fractional_data float32 = 12.022325652356236652312152365223525
	// Prince = 1

	var decided bool = true

	decided = false

	// const confirm = 32
	// confirm = 30

	confirm := 2

	fmt.Println(confirm)

	fmt.Println("My decision on the goa plan is :", decided)
	fmt.Printf("version is %s\n", version)
	fmt.Println("I have money :", money)
	fmt.Println("the interest that i got on my money is ", fractional_data)
}
