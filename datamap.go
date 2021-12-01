package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "Ivan",
		"address": "Bekasi",
	}

	fmt.Println(person["name"])

	var book = make(map[string]string)
	book["tittle"] = "Belajar golang"
	book["author"] = "Ivan"
	book["salah"] = "delete"
	fmt.Println(book)

	delete(book, "salah")
	fmt.Println(book)
}
