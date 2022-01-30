package main

import "fmt"

func main() {

	x := 15
	y:= 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d", x, y)
	}else {
		fmt.Printf("%d is less than %d", y, x)
	}

	color := "blue"
	if color == "red" {
		fmt.Println("\nColor is red")
	}else if color == "blue" {
		fmt.Println("\nColor is blue")
	}else {
		fmt.Println("\nColor is neither blue nor red")
	}

	switch color {
		case "red":
			fmt.Println("Color is red")
		case "blue":
			fmt.Println("Color is blue")
		default :
			fmt.Println("Color is neither red nor blue")
	}

}