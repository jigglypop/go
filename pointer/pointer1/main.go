package main

import "fmt"

type Data struct {
	value int
	data int
}

func main() {
	data := Data{
		999,
		999,
	}
	var p *Data
	p = &data
	fmt.Println(*p)
	fmt.Println(*&data)

}