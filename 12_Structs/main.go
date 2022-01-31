package main

import (
	"fmt"
	"strconv"
)

//Define Person Struct

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (Value Reciever)
func (p Person) greet() string {
	return "Hello, my name is "+p.firstName+" "+p.lastName+" and I am "+strconv.Itoa(p.age)
}


//hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender =="m" {
		return
	}	else
	{
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{"John", "Doe", "Boston", "m", 25}

	fmt.Println(person1)
	person1.age++
	fmt.Println(person1)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2 := Person{"Samantha", "Jones", "Chicago", "f", 19}
	person1.getMarried(person2.lastName)
	fmt.Println(person2.greet())
	person2.getMarried(person1.lastName)

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}