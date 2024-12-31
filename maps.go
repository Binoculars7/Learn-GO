package main

import "fmt"

func runPhoneBook(phoneb map[int]string) {
	for num, val := range phoneb {
		fmt.Printf("The number: %v, belongs to %v \n", num, val)
	}
}

func main() {
	menu := map[string]float64{
		"rice":  45.55,
		"beans": 60.932,
		"garri": 409.12,
	}

	fmt.Println(menu)
	fmt.Println(menu["rice"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		9056070710: "bola",
		8054332839: "funke",
		7012345678: "jacob",
	}

	runPhoneBook(phonebook)

	path := [3]string{"water", "fish", "dog"}

	paths := []string{"water", "fish", "dog", "milk", "cat"}

	fmt.Println(path, paths)

}
