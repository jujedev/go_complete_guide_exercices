package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	/*Si los atributos empiezan con mayuscula, permiten editar externo a la struct (publicos)
	FirstName*/
	firstName string
	/*En cambio si el atributo comienza con minuscula, solo se podría editar mediante algun metodo (privado)*/
	lastName  string
	bithDate  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			bithDate:  "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are require")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		bithDate:  birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.bithDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
