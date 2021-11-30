package main

import "fmt"

func main() {
	var name string
	var friendname = "Budi"
	var age int8

	name = "Ivan Hermawan"
	fmt.Println(name)

	fmt.Println(friendname)

	age = 37
	fmt.Println(age)

	//ganti var dengan :
	//declare pertama
	country := "Indonesia"
	fmt.Println(country)

	//declare kedua
	country = "America"
	fmt.Println(country)

	//var multiple
	var (
		firstname = "Ivan"
		lastname  = "Hermawan"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

	//tipe data constant
	//const kota = "Bekasi"
	//const provinsi = "Jawa Barat"
	//const nomorrumah = 145
	const (
		kota       = "Bekasi"
		provinsi   = "Jawa Barat"
		nomorrumah = 145
	)

	fmt.Println(kota)

}
