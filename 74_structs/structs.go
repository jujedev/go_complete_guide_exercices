package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	bithDate  string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		bithDate:  birthdate,
		createdAt: time.Now(),
	}
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.bithDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirtDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	appUser = newUser(userFirstName, userLastName, userBirtDate)
	// ... do something awesome with that gathered data!
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
