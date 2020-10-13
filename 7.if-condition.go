package main

import "fmt"

func main() {
	var point = 3

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("Tidak lulus. Nilai anda %d hehe", point)
	}
}