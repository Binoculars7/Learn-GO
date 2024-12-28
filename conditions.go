package main

import "fmt"

func main() {

	// age := 30

	// fmt.Println(age <= 20)
	// fmt.Println(age == 30)
	// fmt.Println(age != 30)

	// if age > 40 {
	// 	fmt.Println("age is less that 30")
	// } else if age > 50 {
	// 	fmt.Println("age is less that 30")
	// } else {
	// 	fmt.Println("age is 30")
	// }

	names := []string{"ja", "ba", "co", "so", "me", "you"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at", names[1])
			continue
		}
		if index > 3 {
			fmt.Println("stopping at", index)
			break
		}
		fmt.Println(value)
	}

}
