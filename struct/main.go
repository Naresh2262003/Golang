package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House_No int
	Area     string
	State    string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	// 1st Person
	var prince Person
	prince.FirstName = "Naresh"
	prince.LastName = "Singh"
	prince.Age = 22
	// fmt.Println("Prince object is :", prince)

	// 2nd Method
	// var karan = Person{
	// 	FirstName: "Karan",
	// 	LastName:  "Singh",
	// 	Age:       16,
	// }
	// fmt.Println("Karan object is :", karan)

	// Aarti := Person{
	// 	FirstName: "Aarti",
	// 	LastName:  "Singh",
	// 	Age:       25,
	// }
	// fmt.Println("Aarti Object is ", Aarti)

	// 3rd Person (new Keyword)
	Geeta := new(Person)
	Geeta.FirstName = "Geeta"
	Geeta.LastName = "Devi"
	Geeta.Age = 45

	fmt.Println(*Geeta)

	var kash Employee
	kash.Person_Details = Person{
		FirstName: "kash",
		LastName:  "Singh",
		Age:       22,
	}

	kash.Person_Address = Address{
		House_No: 2,
		Area:     "Uttam Nagar",
		State:    "New Delhi",
	}

	kash.Person_Contact = Contact{
		Email: "Naresh@gmail",
		Phone: "9536856322",
	}

	fmt.Println(kash)

	krish := Employee{
		Person_Details: Person{
			FirstName: "Krish",
			LastName:  "Singh",
			Age:       22,
		},
		Person_Contact: Contact{
			Email: "Maresh",
			Phone: "203",
		},
		Person_Address: Address{
			House_No: 22,
			Area:     "Bikas Po",
			State:    "Momki",
		},
	}
	fmt.Println(krish)

	trish := new(Employee)

	trish.Person_Address = Address{
		House_No: 22,
		Area:     "Bikas Po",
		State:    "Momki",
	}
	trish.Person_Contact = Contact{
		Email: "Maresh",
		Phone: "203",
	}
	trish.Person_Details = Person{
		FirstName: "Trish",
		LastName:  "Singh",
		Age:       22,
	}

	fmt.Println(*trish)

}
