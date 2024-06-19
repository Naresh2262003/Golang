package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)
	studentGrades["Prince"] = 25
	studentGrades["Naresh"] = 23
	studentGrades["Kamlesh"] = 52
	studentGrades["Aryan"] = 62
	studentGrades["Drone"] = 32
	studentGrades["Riyaz"] = 69

	// fmt.Println("Naresh's marks :", studentGrades["Naresh"])
	// delete(studentGrades, "Drone")

	// marks, isPresent := studentGrades["Drone"]
	// fmt.Printf("Drone's marks :%d and isPresent is %t", marks, isPresent)

	// looping over maps
	for key, value := range studentGrades {
		fmt.Printf("%s marks's : %d \n", key, value)
	}

	person := map[string]int{
		"class":   2,
		"roll_no": 21573,
	}

	for key, value := range person {
		fmt.Println("key:", key, "value:", value)
	}

}
