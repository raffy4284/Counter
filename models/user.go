package models

type Email string

type User struct {
	id          int
	firstName   int
	lastName    int
	email       Email
	timeCreated int
	timeUpdated int
}

func (user User) getCounts() []Counter {
	var counters []Counter
	counters = append(counters, Counter{1, "count1", 5, &user, 2, 2})
	counters = append(counters, Counter{2, "count2", 5, &user, 2, 2})
	counters = append(counters, Counter{3, "count3", 5, &user, 2, 2})
	return counters
}

/*
type Counter struct {
	ID          int
	name        int
	count       int
	user        *User
	timeCreated int
	timeUpdated int
}
*/
