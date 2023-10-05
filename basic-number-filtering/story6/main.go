package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bouroo/one2n-go-bootcamp/basic-number-filtering/utils"
)

func main() {
	// Prompt the user for input
	fmt.Print("Sample Input: ")

	// Read the input from the user
	input := bufio.NewScanner(os.Stdin)

	// Read the user's input
	input.Scan()

	// Extract the numbers from the input
	numbers := utils.ExtractNumbers(input.Text())

	// Filter out odd numbers
	result := utils.Filter(numbers, utils.IsOdd)

	// Filter out numbers that are not multiples of three
	result = utils.Filter(result, IsMultipleByThree)

	// Filter out numbers that are not greater than ten
	result = utils.Filter(result, IsGreatherThanTen)

	// Print the output
	fmt.Printf("Sample Output: %s\n", utils.OutputString(result))
}

// IsMultipleByThree checks if a given number is a multiple of three.
//
// Parameters:
// - number: an integer number to be checked.
//
// Returns:
// - a boolean value indicating whether the number is a multiple of three.
func IsMultipleByThree(number int) bool {
	return number%3 == 0
}

// IsGreatherThanTen checks if the given number is greater than ten.
//
// number: an integer to be checked.
// returns: a boolean indicating if the number is greater than ten.
func IsGreatherThanTen(number int) bool {
	return number > 10
}
