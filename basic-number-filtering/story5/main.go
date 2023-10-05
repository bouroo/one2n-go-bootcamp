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

	// Create a new scanner to read user input from the standard input
	input := bufio.NewScanner(os.Stdin)

	// Read the user's input
	input.Scan()

	// Extract numbers from the user input
	numbers := utils.ExtractNumbers(input.Text())

	// Filter the numbers to keep only the even ones
	result := utils.Filter(numbers, utils.IsEven)

	// Filter the result again to keep only the numbers that are multiples of five
	result = utils.Filter(result, IsMultipleByFive)

	// Print the sample output
	fmt.Printf("Sample Output: %s\n", utils.OutputString(result))
}

// IsMultipleByFive checks if the given number is a multiple of five.
//
// Parameters:
//   number - an integer to check
//
// Returns:
//   bool - true if the number is a multiple of five, false otherwise.
func IsMultipleByFive(number int) bool {
	return number%5 == 0
}
