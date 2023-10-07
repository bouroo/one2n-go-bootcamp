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
	// Prompt the user to enter input
	fmt.Print("Sample Input: ")

	// Create a scanner to read user input from the input reader
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()

	// Extract integers from the user input
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter the numbers to get only the prime numbers
	result := utils.Filter(numbers, utils.IsPrime)

	// Filter the result to get only the odd numbers
	result = utils.Filter(result, utils.IsOdd)

	// Create a writer to write the output to the output writer
	outputWriter := bufio.NewWriter(output)

	// Ensure that the writer is flushed at the end
	defer outputWriter.Flush()

	// Write the output string to the output writer
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
