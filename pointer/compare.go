package main

import "fmt"

func main() {
	a := 10
	b := 20

	p1 := &a
	p2 := &a
	p3 := &b

	fmt.Printf("p1의 값 : %d\n", p1)
	fmt.Printf("p2의 값 : %d\n", p2)
	fmt.Printf("p3의 값 : %d\n", p3)

}
