package main

import (
	"fmt"
	"sort"
)

func main() {
	//greetings := "Season greetings to you all"

	// fmt.Println(strings.Contains(greetings, "you"))
	// fmt.Println(strings.ReplaceAll(greetings, "greetings", "messages"))
	// fmt.Println(strings.ToUpper(greetings))
	//fmt.Println(strings.Index(greetings, "ll"))
	//fmt.Println(strings.Split(greetings, " "))

	ages := []int{20, 40, 50, 59, 29, 58, 70, 80, 90}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 59)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "bella", "luka"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "mario"))
}
