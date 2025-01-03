package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func takeBill() bill {
	reader := bufio.NewReader(os.Stdin)
	linput, _ := getInput("Enter the billing type: ", reader)

	mybiller := biller(linput)

	return mybiller

}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Select Option (a - add items, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {

	yourbill := takeBill()
	promptOption(yourbill)
	//fmt.Println(yourbill)

}
