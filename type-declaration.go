package main

import "fmt"

func main() {
	type NoKtp int8
	type NamaKtp string
	type Merried bool

	var NoKtpIvan = 123456789
	var NamaKtpIvan = "Ivan Hermawan"
	var MerriedStatus = false
	fmt.Println(NoKtpIvan)
	fmt.Println(NamaKtpIvan)
	fmt.Println(MerriedStatus)
}
