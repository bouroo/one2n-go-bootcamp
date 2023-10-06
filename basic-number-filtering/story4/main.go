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

	// Filter the numbers by keeping only the prime ones using a utility function
	result := utils.Filter(numbers, utils.IsPrime)
	// Filter the numbers by keeping only the prime odd ones using a utility function
	result = utils.Filter(result, utils.IsOdd)

	// Create a buffered writer to write the output
	outputWriter := bufio.NewWriter(output)

	// Flush the buffered writer to ensure all data is written to the output
	defer outputWriter.Flush()

	// Write the output header and output string
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
