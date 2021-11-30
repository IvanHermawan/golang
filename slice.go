package main

import "fmt"

func main() {
	var months = []string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Di edit"
	fmt.Println(slice1)

	slice1[1] = "Juni lagi"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	var slice3 = append(slice2, "Ivan")
	fmt.Println(slice3)

	slice3[1] = "Desember edit"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	//make slice
	var newSlice = make([]string, 2, 10)

	newSlice[0] = "Ivan"
	newSlice[1] = "Hermawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	var copySLice = make([]string, len(newSlice), cap(newSlice))
	copy(copySLice, newSlice)

	fmt.Println(copySLice)

	var iniArray = [5]int{1, 2, 3, 4, 5}
	var iniSlice = []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
