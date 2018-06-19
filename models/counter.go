package models

type Counter struct {
	ID          int
	name        string
	count       int
	user        *User
	timeCreated int
	timeUpdated int
}
