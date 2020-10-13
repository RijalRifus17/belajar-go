package main

import "fmt"

func main() {
	var decimalNumber = 2.62

	fmt.Printf("Bilangan desimal : %.3f Oke \n", decimalNumber)

	// boolean
	exist := true

	fmt.Printf("exist? %t \n", exist)

	// string
	var message = `Nama saya "Rijal Solahudin".
	Salam kenal.
	Mari Belajar "Golang"`

	fmt.Printf(message)
}