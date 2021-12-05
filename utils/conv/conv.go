package conv

import (
	"strconv"
	"strings"
)

func ToStringSlice(slice []int) []string {
	sliceToReturn := []string{}

	for _, current := range slice {
		convertedInt := strconv.Itoa(current)
		sliceToReturn = append(sliceToReturn, convertedInt)
	}

	return sliceToReturn
}

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		current = strings.ReplaceAll(current, " ", "")
		if len(current) == 0 {
			continue
		}

		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic(err)
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}

func ToInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}

func ToIntOrElse(str string, elseVal int) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		return elseVal
	}

	return number
}
