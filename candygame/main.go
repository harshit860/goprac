package main

import "fmt"

func main() {
	candys := makeCandy()
	eating(candys)
}

func eating(c candys) {
	if len(c) == 0 {
		fmt.Println("Burp....  finished")
	} else {
		fmt.Println("eating..")
		c = c[0 : len(c)-1]
		fmt.Println(c)
		eating(c)
	}
}
