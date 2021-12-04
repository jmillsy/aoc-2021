package main

import (
	"aoc-2021/utils/conv"
	"aoc-2021/utils/files"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(3, "\n")

	powerConsumptionSolution, err := parsePowerConsumption(input) // Part 1
	lifeSupportSolution, err := lifeSupportRating(input)          // Part 2

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Power Consumption (Part 1):")
	fmt.Println(powerConsumptionSolution)

	fmt.Println("Life Support Rating (Part 2):")
	fmt.Println(lifeSupportSolution)
}

func parsePowerConsumption(input []string) (int64, error) {

	var gamma strings.Builder
	var epsilon strings.Builder
	for bitPosition := 0; bitPosition < 12; bitPosition++ {

		sum := 0
		for i := 0; i < len(input); i++ {
			char := strings.Split(input[i], "")
			sum += conv.ToInt(char[bitPosition])
		}

		gamma.WriteString(boolToStr(sum > len(input)/2))
		epsilon.WriteString(boolToStr(sum < len(input)/2))
	}

	gammaDecimal, err := strconv.ParseInt(gamma.String(), 2, 64)
	epsilonDecimal, err := strconv.ParseInt(epsilon.String(), 2, 64)

	// fmt.Printf("Gamma=\t\t %s as decimal=%d \n", gamma.String(), gammaDecimal)
	// fmt.Printf("Epsilon=\t %s as decimal=%d \n", epsilon.String(), epsilonDecimal)

	powerConsumption := gammaDecimal * epsilonDecimal

	return powerConsumption, err
}

func boolToStr(val bool) string {
	if val {
		return "1"
	} else {
		return "0"
	}
}

func lifeSupportRating(input []string) (int64, error) {

	oxy := reducer(input, true)
	c02 := reducer(input, false)

	oxyDecimal, err := strconv.ParseInt(oxy[0], 2, 64)
	c02Decimal, err := strconv.ParseInt(c02[0], 2, 64)

	// fmt.Printf("oxy=%s as dec=%d \n", oxy[0], oxyDecimal)
	// fmt.Printf("c02=%s as dec=%d \n", c02[0], c02Decimal)

	return oxyDecimal * c02Decimal, err
}

func reducer(input []string, majorityRules bool) []string {

	for bitPosition := 0; bitPosition < 12; bitPosition++ {

		if len(input) == 1 {
			return input
		}

		sum := 0
		var startsWith0 []string
		var startsWith1 []string
		for i := 0; i < len(input); i++ {

			char := strings.Split(input[i], "")

			sum += conv.ToInt(char[bitPosition])

			if char[bitPosition] == "0" {
				startsWith0 = append(startsWith0, input[i])
			}

			if char[bitPosition] == "1" {
				startsWith1 = append(startsWith1, input[i])
			}
		}

		var majorityOfOnes bool
		if majorityRules {
			majorityOfOnes = float64(sum) >= float64(len(input))/2
		} else {
			majorityOfOnes = float64(sum) < float64(len(input))/2
		}

		if majorityOfOnes {
			input = startsWith1
		} else {
			input = startsWith0
		}
	}

	return input
}
