package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"bola", "james", "kunle", "dolapo"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices 000

	scores := []int{100, 30, 50}
	sizeOfScores := len(scores)
	fmt.Println(scores, sizeOfScores)

	scores[2] = 58
	fmt.Println(scores, sizeOfScores)

	scores = append(scores, 70)
	sizeOfScores = len(scores)
	fmt.Println(scores, sizeOfScores)

	//slice range ()hhhh

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koppa")
	fmt.Println(rangeOne)
}
