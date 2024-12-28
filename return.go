package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fm1, fm2 := getInitials("adebowale esther")
	fmt.Println(fm1, fm2)

	fm3, fm4 := getInitials("julius krob")
	fmt.Println(fm3, fm4)

	fm5, fm6 := getInitials("caezar")
	fmt.Println(fm5, fm6)

	fm7, fm8 := getInitials("jordan")
	fmt.Println(fm7, fm8)
}
