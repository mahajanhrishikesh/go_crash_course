package main

import "fmt"

func main() {

	emails := make(map[string]string)

	//emails["Bob"] = "bob@gmail.com"
	//emails["Sharon"] = "sharon@gmail.com"
	//emails["Mike"] = "mike@gmail.com"

	

	namesArr := [4]string{"Bob", "Nicole", "Monica", "Chandler"}
	for i:=0; i<4; i++ {
		emails[namesArr[i]] = namesArr[i]+"@gmail.com"
	}

	fmt.Println(emails)

	delete(emails, "Bob")

	fmt.Println(emails)
}