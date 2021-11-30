package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 10
		c = a * b
	)

	fmt.Println(c)

	var i = 10
	i += 10

	fmt.Println("Nilai dari i", i)

	i++
	fmt.Println("Nilai dari i++", i)

	//operasi perbandingan
	var (
		name1       = "Ivan"
		name2       = "Ivan"
		comp   bool = name1 != name2
		value1      = 100
		value2      = 200
	)

	fmt.Println(comp)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

	//operasi boolean
	var (
		ujian        = 90
		absensi      = 80
		lulusUjian   = ujian >= 80
		lulusAbsensi = absensi >= 80
		lulus        = lulusUjian && lulusAbsensi
	)

	fmt.Println("Lulus Ujian", lulusUjian)
	fmt.Println("Lulus Absensi", lulusAbsensi)
	fmt.Println("Hasilnya", lulus)
	//fmt.Println("jadi",ujian >= 80 && absensi >= 80)
}
