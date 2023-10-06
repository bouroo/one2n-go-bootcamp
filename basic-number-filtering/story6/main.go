package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/bouroo/one2n-go-bootcamp/basic-number-filtering/utils"
)

func main() {
	MainIO(os.Stdin, os.Stdout)
}

func MainIO(input io.Reader, output io.Writer) {
	// Prompt the user for input
	fmt.Print("Sample Input: ")

	// Create a new scanner to read input from the user
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()

	// Extract the numbers from the input using a utility function
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter the numbers by keeping only the odd ones using a utility function
	result := utils.Filter(numbers, utils.IsOdd)

	// Filter the numbers by keeping only the multiples of 3 ones using a utility function
	result = utils.Filter(result, IsMultipleByThree)

	// Filter the numbers by keeping only the numbers greater than 10 using a utility function
	result = utils.Filter(result, IsGreatherThanTen)

	// Create a buffered writer to write the output
	outputWriter := bufio.NewWriter(output)

	// Flush the buffered writer to ensure all data is written to the output
	defer outputWriter.Flush()

	// Write the output header and output string
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
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
