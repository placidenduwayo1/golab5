package generic2

import (
	"bufio"
	"fmt"
	"os"
)

type Customer struct {
	Firstname string
	Lastname  string
	Email     string
	Birthdate string
}

func NewCustomer(fn, ln, em, bt string) *Customer {
	return &Customer{
		Firstname: fn,
		Lastname:  ln,
		Email:     em,
		Birthdate: bt,
	}
}

func (c *Customer) String() string {
	return fmt.Sprintf("{\n"+
		"\tFirstname: %s\n"+
		"\tLastname: %s\n"+
		"\tEmail: %s\n"+
		"\tBirthdate: %s\n"+
		"}", c.Firstname, c.Lastname, c.Email, c.Birthdate)
}

func CustomerInput() Customer {
	var streamReader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your first name:")
	firstName, _ := streamReader.ReadString('\n')
	fmt.Print("Enter your last name:")
	lastName, _ := streamReader.ReadString('\n')
	fmt.Print("Enter your email:")
	email, _ := streamReader.ReadString('\n')
	fmt.Print("Enter your birthdate:")
	birthdate, _ := streamReader.ReadString('\n')
	return Customer{firstName, lastName, email, birthdate}
}
