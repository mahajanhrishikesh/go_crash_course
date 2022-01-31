package main 

import "fmt"

func main() {
	ids := []int{33,76,54,23,11,2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum:=0
	for _, id := range ids {
		sum+=id
	}
	fmt.Println("Sum ", sum)

	emails := map[string]string{"Bob": "bob@gmail.com", "Jessica": "jessica@gmail.com", "John":"john@usc.edu"}

	// Range with maps
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}