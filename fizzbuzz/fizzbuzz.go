package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(number int) string {
	if isMultipleOfThree(number) {
		return "Fizz"
	}
	if isMultipleOfFive(number) {
		return "Buzz"
	}
	return strconv.Itoa(number)
}

func isMultipleOfFive(number int) bool {
	return number%5 == 0
}

func isMultipleOfThree(number int) bool {
	return number%3 == 0
}
