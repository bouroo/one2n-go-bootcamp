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
	// Split the input string by comma and store the resulting substrings in the 'nums' slice.
	nums := strings.Split(input, ",")

	// Iterate over each substring in the 'nums' slice.
	for _, num := range nums {
		// Trim any leading or trailing whitespace from the current substring.
		trimmedNum := strings.TrimSpace(num)

		// Convert the trimmed substring to an integer.
		// If the conversion is successful, append the number to the 'numbers' slice.
		if number, err := strconv.Atoi(trimmedNum); err == nil {
			numbers = append(numbers, number)
		}
	}

	// Return the 'numbers' slice containing the extracted numbers.
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
	// Check if the number is less than 2, which is not a prime number.
	if number < 2 {
		return false
	}

	// Iterate from 2 to the square root of the number.
	for i := 2; i*i <= int(number); i++ {
		// Check if the number is divisible by any value between 2 and the square root of the number.
		if int(number)%i == 0 {
			return false
		}
	}

	// If the number is not divisible by any value between 2 and the square root of the number, it is a prime number.
	return true
}

// Filter filters the given numbers using the provided filter function.
//
// numbers: The list of numbers to be filtered.
// filterFn: The filter function that determines whether a number should be included in the filtered list.
// filteredNumbers: The list of numbers that pass the filter.
func Filter(numbers []int, filterFn FilterFn) (filteredNumbers []int) {
	// Iterate over each number in the numbers list.
	for _, number := range numbers {
		// Check if the number passes the filter function.
		if filterFn(number) {
			// If the number passes the filter, add it to the filteredNumbers list.
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
	// Create an empty slice to store the string representations of the numbers
	output := []string{}

	// Iterate over each number in the input slice
	for _, number := range numbers {
		// Convert the number to a string and append it to the output slice
		output = append(output, fmt.Sprint(number))
	}

	// Join all the string representations of the numbers with a comma and a space
	result := strings.Join(output, ", ")

	// Return the final result as a string
	return result
}
