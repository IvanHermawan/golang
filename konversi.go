package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 129
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)

		name    = "Ivan Hermawan"
		e       = name[6]
		eString = string(e)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
