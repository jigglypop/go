package main

import "fmt"

func main() {
	var t = [3]string {"monday", "tuesday", "wednesday"}
	for i := 0;i < 3;i++ {
		fmt.Println(t[i])
	}

	x := [...]int {10, 20, 30}
	for i, v := range x {
		fmt.Println(i, v)
	}
}