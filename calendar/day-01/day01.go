package main

import (
	"aoc-2021/utils/conv"
	"aoc-2021/utils/files"
	"fmt"
	"log"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := conv.ToIntSlice(inputSliceAsString)

	// solution, err := depthChange(input) // Part 1
	solution, err := depthChangeSlidingWindow(input) // Part 2

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}

func depthChangeSlidingWindow(input []int) (int, error) {

	previousWindowSum := 0
	currentWindowSum := 0

	increaseCount := 0

	for i := 0; i < len(input); i++ {

		currentWindowSum = 0
		// Peek ahead at the next three
		for j := 0; j < 3; j++ {

			// At the end, skip if there isnt a whole window of 3
			if i+j >= len(input) {
				break
			}

			currentWindowSum += input[i+j]
		}

		if currentWindowSum > previousWindowSum && i != 0 {
			increaseCount++
		}

		previousWindowSum = currentWindowSum

	}

	return increaseCount, nil
}

func depthChange(input []int) (int, error) {

	previous := 0
	current := 0
	increaseCount := 0
	for i := 0; i < len(input); i++ {

		if i == 0 {
			fmt.Println("Skipping first iteration")
			continue
		}

		current = input[i]
		previous = input[i-1]

		if current > previous {
			increaseCount++
			fmt.Printf("INC iteration=%d, p=%d, c=%d \n", i, previous, current)
		} else {
			fmt.Printf("DEC iteration=%d, p=%d, c=%d \n", i, previous, current)
		}
	}

	return increaseCount, nil
}
