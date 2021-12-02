package main

import (
	"aoc-2021/utils/conv"
	"aoc-2021/utils/files"
	"fmt"
	"log"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")

	solution, err := spaceChangesPart1(input) // Part 1
	//solution, err := spaceChangesPart2(input) // Part 2

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}

func spaceChangesPart1(input []string) (int, error) {

	depth := 0
	horizontal := 0

	for i := 0; i < len(input); i++ {
		direction := strings.Split(input[i], " ")[0]
		directionAmount := conv.ToInt(strings.Split(input[i], " ")[1])

		if direction == "forward" {
			horizontal += directionAmount
		}

		if direction == "up" {
			depth -= directionAmount
		}

		if direction == "down" {
			depth += directionAmount
		}

		fmt.Printf("%s\t [%d]\t horizontal=%d \t depth=%d \n", direction, directionAmount, horizontal, depth)
	}

	return depth * horizontal, nil
}

func spaceChangesPart2(input []string) (int, error) {

	depth := 0
	horizontal := 0
	aim := 0

	for i := 0; i < len(input); i++ {
		direction := strings.Split(input[i], " ")[0]
		directionAmount := conv.ToInt(strings.Split(input[i], " ")[1])

		if direction == "forward" {
			horizontal += directionAmount
			depth += aim * directionAmount
		}

		if direction == "up" {
			aim -= directionAmount
		}

		if direction == "down" {
			aim += directionAmount
		}

		fmt.Printf("%s\t [%d]\t horizontal=%d\t depth=%d\t aim=%d \n", direction, directionAmount, horizontal, depth, aim)
	}

	return depth * horizontal, nil
}
