package main

import "fmt"

func main() {
	//var fruitArr [2]string

	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])


	fruitSlice:= []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice[0:2])
}