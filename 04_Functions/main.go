package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("04_Functions")
	fmt.Println(greeting("Nancy"))
	fmt.Println(getSum(300,120))
}