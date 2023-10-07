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
	// Print the prompt for sample input
	fmt.Print("Sample Input: ")
	
	// Read the input from the reader
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()

	// Extract integers from the input text
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter the numbers to keep only the even ones
	result := utils.Filter(numbers, utils.IsEven)

	// Filter the result further to keep only the numbers that are multiples of 5
	result = utils.Filter(result, utils.IsMultipleOf[int](5))

	// Create a buffered writer for the output
	outputWriter := bufio.NewWriter(output)

	// Make sure the writer is flushed after the function returns
	defer outputWriter.Flush()

	// Write the sample output to the writer
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
