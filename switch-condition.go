package main

import "fmt"

func main() {
	var name = "Joko"

	switch name {
	case "Ivan":
		fmt.Println("Ini Ivan")
	case "Joko":
		fmt.Println("Ini Joko")
	default:
		fmt.Println("gk nemu")
	}

	//short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Kepanjangan")
	case false:
		fmt.Println("pas udh")
	}

	//tanpa kondisi
	length1 := len(name)

	switch {
	case length1 > 5:
		fmt.Println("Kepanjangan")
	case length1 < 5:
		fmt.Println("pas udh")
	}
}
