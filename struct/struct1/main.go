package main

import "fmt"

type House struct {
	Address string
	Size int
	Price float64
	Type string
}

func main() {
	var house House
	house.Address = "서울시"
	house.Price = 9.8
	house.Size = 28
	house.Type = "아파트"

	fmt.Println("주소: ", house.Address)
	fmt.Println("가격: ", house.Price)
	fmt.Println("크기: ", house.Size)
	fmt.Println("타입: ", house.Type)

	fmt.Println(house)

	house2 := House{
		"대전시",
		20,
		9.6,
		"빌라",
	}
	
	fmt.Println(house2)
}