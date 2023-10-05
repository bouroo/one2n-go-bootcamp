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

	// Create a new scanner to read input from the user
	input := bufio.NewScanner(os.Stdin)

	// Read the user's input
	input.Scan()

	// Extract the numbers from the input using a utility function
	numbers := utils.ExtractNumbers(input.Text())

	// Filter the numbers by keeping only the even ones using a utility function
	result := utils.Filter(numbers, utils.IsEven)

	// Print the filtered numbers as the sample output
	fmt.Printf("Sample Output: %s\n", utils.OutputString(result))
}
