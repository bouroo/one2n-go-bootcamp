package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type FilterFn[T Number] func(number T) bool

// ExtractIntegers extracts numbers from a comma-separated string.
//
// The input parameter is a string that contains comma-separated numbers.
// The function returns a slice of integers.
func ExtractIntegers(input string) (numbers []int) {
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
func IsEven[T Number](number T) bool {
	return math.Mod(float64(number), 2) == 0
}

// IsOdd determines if a given number is odd.
//
// number: an integer number.
// Returns: a boolean value indicating if the number is odd.
func IsOdd[T Number](number T) bool {
	return math.Mod(float64(number), 2) == 1
}

// IsPrime checks if the given number is prime.
//
// number: an integer number to be checked.
// returns: a boolean indicating whether the number is prime or not.
func IsPrime[T Number](number T) bool {
	// Check if the number is less than 2, which is not a prime number.
	if number < 2 {
		return false
	}

	// Iterate from 2 to the square root of the number.
	for i := 2.0; i*i <= float64(number); i++ {
		// Check if the number is divisible by any value between 2 and the square root of the number.
		if math.Mod(float64(number), i) == 0 {
			return false
		}
	}

	// If the number is not divisible by any value between 2 and the square root of the number, it is a prime number.
	return true
}

// IsGreaterThan returns a FilterFn that checks if a given number is greater than a specified value.
//
// Parameters:
// - number: the number to be compared.
// - greaterThan: the value to compare against.
//
// Returns:
// - FilterFn: a function that takes a number and returns a boolean indicating if it is greater than the specified value.
func IsGreaterThan[T Number](greaterThan T) FilterFn[T] {
	return func(number T) bool {
		return number > greaterThan
	}
}

// IsLessThan is a generic function that returns a filter function
// which checks if a given number is less than a provided value.
//
// Parameters:
// - number: the number to be compared
// - lessThan: the value to compare against
//
// Return type:
// - FilterFn[T]: a function that takes a number and returns a boolean
func IsLessThan[T Number](lessThan T) FilterFn[T] {
	return func(number T) bool {
		return number < lessThan
	}
}

// IsMultipleOf returns a FilterFn[int] that checks if a given number is a multiple of the specified multiplyer.
//
// Parameters:
//   - number: an integer representing the number to be checked.
//   - multiplyer: an integer representing the number to check if it is a multiple of.
//
// Returns:
//   - FilterFn[int]: a function that takes an integer as input and returns a boolean indicating if it is a multiple of the multiplyer.
func IsMultipleOf[T Number](multiplyer T) FilterFn[T] {
	return func(number T) bool {
		return math.Mod(float64(number), float64(multiplyer)) == 0
	}
}

// Filter filters the given numbers using the provided filter function.
//
// numbers: The list of numbers to be filtered.
// filterFn: The filter function that determines whether a number should be included in the filtered list.
// filteredNumbers: The list of numbers that pass the filter.
func Filter[T Number](numbers []T, filterFn FilterFn[T]) (filteredNumbers []T) {
	// Iterate over each number in the numbers list.
	for _, number := range numbers {
		// Check if the number passes the filter function.
		if filterFn(T(number)) {
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
	output := strings.Builder{}
	for i, number := range numbers {
		if i > 0 {
			output.WriteString(", ")
		}
		output.WriteString(fmt.Sprint(number))
	}
	return output.String()
}
