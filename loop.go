package main

import (
	"fmt"
	"sort"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is", i)
	// }

	// names := []string{"bola", "james", "tolu", "titi"}
	// sizeOfNames := len(names)
	// for i := 0; i < sizeOfNames; i++ {
	// 	name := names[i]
	// 	fmt.Printf("Numbers %v is %v \n", i, name)
	// }

	names := []string{"ja", "ba", "co", "so", "me", "you"}
	sort.Strings(names)

	for _, value := range names {
		fmt.Printf("%v \n", value)
	}

}
