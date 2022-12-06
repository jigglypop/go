package main

import "fmt"

func main() {
	x := [...]int {10, 20, 30}
	for _, v := range x {
		fmt.Printf("[%d]\n", v)
	}
}