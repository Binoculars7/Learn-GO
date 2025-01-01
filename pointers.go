package main

import "fmt"

func updateVal(x *string) string {
	*x = "cold"
	return *x
}

func main() {
	state := "hot"

	stateMemAddress := &state

	///fmt.Println(state)
	fmt.Println(stateMemAddress)
	fmt.Println(state)
	fmt.Println(updateVal(stateMemAddress))
	fmt.Println(state)
}
