package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Ivan"
	name[1] = "Hermawan"

	fmt.Println(name[1])

	var values = [3]int{
		99,
		98,
		97,
	}

	fmt.Println(values[1])

	fmt.Println(len(name))
}
