package main

import "fmt"

func main() {
	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("None of the above")
	}

	month := "january"

	switch month {
	case "November", "December", "January", "February":
		fmt.Println("Winter")
	case "March", "April":
		fmt.Println("Spring")
	case "May", "June", "July", "August":
		fmt.Println("Summers")
	case "September", "October":
		fmt.Println("Autumn")
	default:
		fmt.Println("Enter the name of the Month in correct format as 'January'")
	}

}
