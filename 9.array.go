package main

import "fmt"

func main ()  {
	var names [4]string 
	names[0] = "Luklu"
	names[1] = "Rijal"
	names[2] = "Aziz"
	names[3] = "Nisa"

	fmt.Println(names)

	/////////////////////////////////

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi Array seluruhnya \t", fruits)


	/////////////////////////////////////

	// cara horizontal
	fruits = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)

	///////////////////////////////////////////
	var numbers = [...]int{1,2,3,4,5}

	///////////////////////////////////////// ARRAY MULTI DEMENSI ///////////////
	var numbers1 = [2][3]int{[3]int{3,2,3}, [3]int{1,2,3}}
	var numbers2 = [2][3]int{{ 3,2,3}, {4,5,5 }}

	fmt.Println(numbers, numbers1, numbers2)


	/////////////////////////////////////////////////////
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits2); i++ {
	// 	fmt.Printf("Element %d: %s\n", i ,fruits2[i])
	// }

	// for i, fruit := range fruits2 {
	// 	fmt.Printf("Element %d: %s\n", i ,fruit)
	// }

	for _, fruit := range fruits2 {
		fmt.Printf("Element : %s\n" ,fruit)
	}

	///////////////////////
	var fruits3 = make([]string, 2)
	fruits3[0] = "Apple"
	fruits3[1] = "Grape"

	fmt.Println(fruits3)

}