package utils

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type FilterFn func(number int) bool

// ExtractNumbers extracts numbers from a comma-separated string.
//
// The input parameter is a string that contains comma-separated numbers.
// The function returns a slice of integers.
func ExtractNumbers(input string) (numbers []int) {
	nums := strings.Split(input, ",")
	for _, num := range nums {
		if number, err := strconv.Atoi(strings.TrimSpace(num)); err == nil {
			numbers = append(numbers, number)
		}
	}
	return
}

// IsEven checks if a given number is even.
//
// number: an integer that needs to be checked.
// bool: true if the number is even, false otherwise.
func IsEven(number int) bool {
	return int(number)%2 == 0
}

// IsOdd determines if a given number is odd.
//
// number: an integer number.
// Returns: a boolean value indicating if the number is odd.
func IsOdd(number int) bool {
	return int(number)%2 != 0
}

// IsPrime checks if the given number is prime.
//
// number: an integer number to be checked.
// returns: a boolean indicating whether the number is prime or not.
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= int(number); i++ {
		if int(number)%i == 0 {
			return false
		}
	}
	return true
}

// Filter filters the given numbers using the provided filter function.
//
// numbers: The list of numbers to be filtered.
// filterFn: The filter function that determines whether a number should be included in the filtered list.
// filteredNumbers: The list of numbers that pass the filter.
func Filter(numbers []int, filterFn FilterFn) (filteredNumbers []int) {
	for _, number := range numbers {
		if filterFn(number) {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return
}

// OutputString generates a string representation of the given numbers.
//
// The numbers parameter is a slice of any numeric type.
// The function returns a string.
func OutputString[T Number](numbers []T) string {
	output := []string{}
	for _, number := range numbers {
		output = append(output, fmt.Sprint(number))
	}
	return strings.Join(output, ", ")
}
