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

// MainIO function takes an input reader and an output writer as parameters.
func MainIO(input io.Reader, output io.Writer) {
	// Prompt user for input
	fmt.Print("Sample Input: ")

	// Create a scanner to read input from the reader
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()

	// Extract integers from the input text
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter out odd numbers from the list
	result := utils.Filter(numbers, utils.IsOdd)

	// Filter out numbers that are not multiples of 3
	result = utils.Filter(result, utils.IsMultipleOf[int](3))

	// Filter out numbers that are not greater than 10
	result = utils.Filter(result, utils.IsGreaterThan[int](10))

	// Create a writer to write output to the output writer
	outputWriter := bufio.NewWriter(output)

	// Make sure the writer is flushed at the end
	defer outputWriter.Flush()

	// Write the sample output to the output writer
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
