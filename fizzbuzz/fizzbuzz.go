package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if isMultipleOfThree(number) {
		return "Fizz"
	}
	return strconv.Itoa(number)
}

func isMultipleOfThree(number int) bool {
	return number%3 == 0
}
