package main

import "fmt"

func main() {
	var point = 3

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna\n")
	} else if point > 5 {
		fmt.Println("lulus\n")
	} else if point == 4 {
		fmt.Println("Hampir Lulus\n")
	} else {
		fmt.Printf("Tidak lulus. Nilai anda %d hehe\n", point)
	}

	var point2 = 4300.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%\n")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%\n")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%\n")
	}

	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("Perfect")
	case 7,6,5,4:
		fmt.Println("Awasome")
	default:
		fmt.Println("Not Bad")
	}
}