package main

import "fmt"

func main() {
	var name = "Jokoooooo"

	if name == "Ivan" {
		fmt.Println("Ini Ivan")
	} else if name == "Joko" {
		fmt.Println("Ini Joko")
	} else {
		fmt.Println("Gk nemu")
	}

	//var length = len(name)
	//short statement
	if length := len(name); length > 5 {
		fmt.Println("kepanjangan")
	} else {
		fmt.Println("sudah benar")
	}

}
