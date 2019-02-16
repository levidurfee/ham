package models

// User contains user information.
// They have to be logged in before they can have any user information
// in this model. So, it might be empty.
type User struct {
	Email string
	UID   string
}
