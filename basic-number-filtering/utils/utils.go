package utils

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type FilterFn[T Number] func(number T) bool

var RegNumPool = sync.Pool{
	New: func() any {

		return regexp.MustCompile(`\d+`)
	},
}

// ExtractIntegers splits the input string by comma to get individual numbers and
// converts them to integers. It returns a slice of integers.
//
// Parameters:
// - input: a string containing comma-separated numbers
//
// Return type:
// - numbers: a slice of integers
func ExtractIntegers(input string) (numbers []int) {
	// Split the input string by comma to get individual numbers
	nums := strings.Split(input, ",")

	// Iterate over each number
	for _, num := range nums {
		// Trim any leading or trailing spaces from the number
		trimmedNum := strings.TrimSpace(num)

		// Convert the trimmed number to an integer
		if number, err := strconv.Atoi(trimmedNum); err == nil {
			// Append the integer to the numbers slice
			numbers = append(numbers, number)
		}
	}

	// Return the numbers slice
	return
}

// IsEven checks if the given number is even.
//
// It takes a parameter 'number' of type T, which should be a number.
// It returns a boolean value indicating whether the number is even or not.
func IsEven[T Number](number T) bool {
	return math.Mod(float64(number), 2) == 0
}

// IsOdd checks if the given number is odd.
//
// - `number`: The number to be checked.
// - Returns `bool`: `true` if the number is odd, `false` otherwise.
func IsOdd[T Number](number T) bool {
	return math.Mod(float64(number), 2) != 0
}

// IsPrime checks if the given number is prime.
//
// The number parameter is the number to be checked.
// The function returns a boolean value indicating whether the number is prime or not.
func IsPrime[T Number](number T) bool {

	// Check if the number is less than 2.
	if number < 2 {
		return false
	}

	// Check divisibility of the number from 2 up to the square root of the number.
	for i := 2.0; i*i <= float64(number); i++ {

		// If the number is divisible by i, it is not prime.
		if math.Mod(float64(number), i) == 0 {
			return false
		}
	}

	// If the number is not divisible by any number up to its square root, it is prime.
	return true
}

// IsGreaterThan returns a FilterFn function that takes a parameter greaterThan of type T and returns a boolean value.
//
// The FilterFn function returned takes a parameter number of type T and checks if the number is greater than greaterThan.
// It returns true if the number is greater than greaterThan, and false otherwise.
func IsGreaterThan[T Number](greaterThan T) FilterFn[T] {
	return func(number T) bool {
		return number > greaterThan
	}
}

// IsLessThan returns a filter function that checks if a given number is less than the specified value.
//
// Parameters:
// - lessThan: The value to compare against.
//
// Return Type:
// - FilterFn[T]: A function that takes a number and returns a boolean indicating if it is less than the specified value.
func IsLessThan[T Number](lessThan T) FilterFn[T] {
	return func(number T) bool {
		return number < lessThan
	}
}

// IsMultipleOf returns a filter function that checks if a given number is a multiple of the provided multiplier.
//
// The `multiplier` parameter is the value by which the number should be a multiple of.
// The function returns a boolean value indicating whether the number is a multiple of the multiplier.
func IsMultipleOf[T Number](multiplyer T) FilterFn[T] {
	return func(number T) bool {
		return math.Mod(float64(number), float64(multiplyer)) == 0
	}
}

// Filter filters a slice of numbers based on a given filter function.
//
// numbers: The slice of numbers to filter.
// filterFn: The filter function to apply to each number.
// filteredNumbers: The filtered slice of numbers.
func Filter[T Number](numbers []T, filterFn FilterFn[T]) (filteredNumbers []T) {
	// Iterate over each number in the input slice
	for _, number := range numbers {
		// Apply the filter function to the current number
		// Convert the number to type T before passing it to the filter function
		if filterFn(T(number)) {
			// If the filter function returns true, add the number to the filtered slice
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return
}

// FilterAll filters the given numbers using the provided filter functions.
//
// numbers: The slice of numbers to filter.
// filterFn: The variadic slice of filter functions to apply.
// filteredNumbers: The slice of filtered numbers.
func FilterAll[T Number](numbers []T, filterFn ...FilterFn[T]) (filteredNumbers []T) {
	// Initialize a boolean variable to keep track of whether a number is valid or not.
	var valid = false

	// Loop through each number in the given slice.
numLoop:
	for _, num := range numbers {
		// Create a temporary variable to store the valid number.
		var tmpNum T

		// Reset the validity flag for each number.
		valid = false

		// Loop through each filter function.
		for _, fn := range filterFn {
			// Check if the number passes the filter function.
			if !fn(num) {
				// If the number fails the filter function, skip to the next number.
				continue numLoop
			}

			// If the number passes the filter function, set the validity flag to true
			// and update the temporary number variable.
			valid = true
			tmpNum = num
		}

		// If the number is valid, append it to the filtered numbers slice.
		if valid {
			filteredNumbers = append(filteredNumbers, tmpNum)
		}
	}

	// Return the filtered numbers slice.
	return
}

// FilterAny filters the given numbers based on the provided filter functions.
//
// The numbers parameter is a slice of any type that satisfies the Number interface.
// The filterFn parameter is a variadic parameter of type FilterFn[T], which is a function type that takes a value of type T and returns a boolean.
// The filteredNumbers parameter is a slice of the same type as the numbers slice.
// The function returns the filteredNumbers slice.
func FilterAny[T Number](numbers []T, filterFn ...FilterFn[T]) (filteredNumbers []T) {
	// Initialize a boolean variable to track if the current number is valid based on the provided filter functions.
	var valid = false
	
	// Iterate over each number in the given slice.
	for _, num := range numbers {
		// Create a temporary variable to store the current number.
		var tmpNum T
		
		// Reset the valid flag to false for each new number.
		valid = false
		
		// Loop through each filter function.
	filterLoop:
		for _, fn := range filterFn {
			// Check if the current number satisfies the filter function.
			if fn(num) {
				// Set the valid flag to true.
				valid = true
				
				// Store the current number in the temporary variable.
				tmpNum = num
				
				// Exit the filter loop since a match has been found.
				break filterLoop
			}
		}
		
		// If the current number is valid, append it to the filteredNumbers slice.
		if valid {
			filteredNumbers = append(filteredNumbers, tmpNum)
		}
	}
	
	// Return the filteredNumbers slice.
	return
}

// OutputString concatenates the elements of a []Number into a single string.
//
// numbers: The []Number to be concatenated.
// string: The concatenated string.
func OutputString[T Number](numbers []T) string {
	// Create a strings.Builder to efficiently build the output string
	output := strings.Builder{}
	
	// Iterate over each element in the numbers slice
	for i, number := range numbers {
		// If it's not the first element, add a comma and space to the output string
		if i > 0 {
			output.WriteString(", ")
		}
		
		// Convert the number to a string and add it to the output string
		output.WriteString(fmt.Sprint(number))
	}
	
	// Return the final concatenated string
	return output.String()
}
