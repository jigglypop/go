package main

import "fmt"

func main() {
	slice1 := []string{"A", "B", "C"}
	for i := 0;i < len(slice1);i++ {
		println(slice1[i])
	}

	for i, v := range slice1 {
		println(i, v)
	}
	// append

	slice := []int{1, 2, 3}
	slice2 := append(slice, 4)
	fmt.Println(slice2)

	slice = append(slice, 3, 4, 5, 6, 7)
	fmt.Println(slice)
}
