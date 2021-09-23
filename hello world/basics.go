package main

import "fmt"

func main() {
	// var name string = "harshit"

	numbers := helper1{1, 2, 3}
	// array is fixed length in array
	// slice is like normal array in go lang
	numbers.print()
	basicsDeclare()
}

func fullName() int {
	return 2
}

func basicsDeclare() {

	var name1 string = "harshit" // static declaration
	fmt.Println(name1)
	name := fullName() // dynamic declaration
	fmt.Println(name)

}
