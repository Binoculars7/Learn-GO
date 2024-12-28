package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [2]string{"bola", "james"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices

	scores := []int{100, 30, 50}
	sizeOfScores := len(scores)
	fmt.Println(scores, sizeOfScores)

	scores[2] = 58
	fmt.Println(scores, sizeOfScores)

	scores = append(scores, 70)
	sizeOfScores = len(scores)
	fmt.Println(scores, sizeOfScores)
}
