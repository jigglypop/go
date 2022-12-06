package main

import "fmt"

func main() {
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var two [2][3]int
	for y := 0; y < 2; y++ {
		for x := 0; x < 3; x++ {
			two[y][x] = y + x
		}
	}
	fmt.Println(two)
}
