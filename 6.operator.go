package main

import "fmt"

func main() {
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrLeft = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrLeft)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}