package models

import (
	database "github.com/raffy4284/counter/models/database"
)

type Email string

type User struct {
	id          int
	firstName   string
	lastName    string
	email       Email
	timeCreated int
	timeUpdated int
}

func CreateUser(firstName string, lastName string, email Email) User {
	user := User{
		firstName: firstName,
		lastName:  lastName,
		email:     email}
	database.GetDatabase().DB.Create(&user)
	return user
}

func DeleteUser(id int) {
	database.GetDatabase().DB.Delete(User{}, "id LIKE ?", "%"+string(id)+"%")
}

// TODO: replace with actually using ORM
func (user User) getCounts() []Counter {
	var counters []Counter
	counters = append(counters, Counter{1, "count1", 5, &user, 2, 2})
	counters = append(counters, Counter{2, "count2", 5, &user, 2, 2})
	counters = append(counters, Counter{3, "count3", 5, &user, 2, 2})
	return counters
}
