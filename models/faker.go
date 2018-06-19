package models

type Person struct {
	ID        string
	Firstname string
	Lastname  string
	Address   *Address
}

type Address struct {
	City  string
	State string
}

func GetPeople() []Person {
	var people []Person
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	return people
}
