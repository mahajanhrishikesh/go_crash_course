package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Println("%T\n",a)
	fmt.Println("%T\n", b)

	// Use * to read value at address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Everything in go is pass by value
}