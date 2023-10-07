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

	// Create a scanner to read the input
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()

	// Extract integers from the input string
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter the numbers to keep only the odd ones
	result := utils.Filter(numbers, utils.IsOdd)

	// Create a buffered writer to write the output
	outputWriter := bufio.NewWriter(output)

	// Make sure the writer is flushed before the function exits
	defer outputWriter.Flush()

	// Write the sample output to the writer
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
