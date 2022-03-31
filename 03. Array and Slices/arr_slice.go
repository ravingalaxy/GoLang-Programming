package main

import (
	"fmt"
)

func main() {
	var firstNames = [3]string{"Ravin", "Ravindra", "Sidh"}
	fmt.Println(firstNames[1])

	var nummies = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nummies[0])

	lastName := [3]string{"Singh", "kumar", "Sagar"}
	fmt.Println(lastName[0])

	ourNames := []string{"elder", "younger"}
	fmt.Println(ourNames)
	ourNames[1] = "Smith"
	fmt.Println(ourNames)

	var ourNummies = [5]int{1, 2}
	fmt.Println(ourNummies)

	var moreNummies = [10]int{0: 41, 4: 99}
	fmt.Println(moreNummies)

	//slices
	var toppings = [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(toppings)

	toppingSlice := toppings[0:2]
	fmt.Println(toppingSlice)

	//modify slice
	toppingSlice[1] = "Peppers"
	fmt.Println(toppingSlice)

	// add item to slice
	toppingSlice = append(toppingSlice, "Apple")
	fmt.Println(toppingSlice)

	//add two slices together
	otherSlice := toppings[3:4]
	fmt.Println(otherSlice)
	otherSlice = append(otherSlice, toppingSlice...)
	fmt.Println(otherSlice)
}
