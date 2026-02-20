package person

import (
	"errors"
	"fmt"
	"time"
)

type person struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type admin struct {
	person
	Level string
}

func New(firstName, lastName, birthdate string) (*person, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &person{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthdate string) (*admin, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &admin{
		person: person{
			firstName: firstName,
			lastName:  lastName,
			birthdate: birthdate,
			createdAt: time.Now(),
		},
		Level: "admin",
	}, nil
}

func (p *person) PrintData() {
	fmt.Println(p.firstName, p.lastName, p.birthdate)
}

func (p *person) UpdateBirthdate(birthdate string) {
	p.birthdate = birthdate
}
