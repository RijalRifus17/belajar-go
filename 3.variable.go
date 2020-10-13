package main

import "fmt"

func main() {
	// GENERAL
	// var firstName string = "Rijal"

	// var lastName string
	// lastName = "Solahudin"

	// fmt.Println("Hallo", firstName, lastName)
	
	// TANPA VAR
	// job := "Programmer"
	// job = "Penguasaha"
	
	// fmt.Println(job)

	// Multi Variable
	// var first, second, third string
	// first, second, third = "Satu", "Dua", "Tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eigth, ninth := "7", "8", "9"

	// undescore
	// _ = "Golang"
	// _ = "Golang itu mudah"

	// name, _ = "Rijal", "Solahudin"

	// Keyword New

	name := new(string)

	fmt.Println(*name)
}