package main

import "fmt"

var names []string = []string{"bol", "james", "wale", "john"}
var i int
var count = len(names)

func sayHi() {
	fmt.Printf("%v is highly welcome to the party \n", names[i])
}

func sayBye() {
	fmt.Printf("%v is not invited to the party \n", names[i])
}

func main() {
	for i = 0; i < count; i++ {
		if i%2 == 0 {
			sayHi()
		} else {
			sayBye()
		}
	}
}
