package main

import "fmt"

type helper1 []int

func (n helper1) print() {
	for index, value := range n {
		fmt.Println(index, value)
	}
}
