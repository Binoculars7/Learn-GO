package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

func biller(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tips:  0,
	}
	return b
}

func main() {
	mybill := biller("Youth Young")

	fmt.Println(mybill)
}
