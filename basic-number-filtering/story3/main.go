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
	// Print the prompt for sample input.
	fmt.Print("Sample Input: ")

	// Create a scanner to read the input.
	inputScanner := bufio.NewScanner(input)
	// Scan the input and store it as a string.
	inputScanner.Scan()

	// Extract the integers from the input string using the `ExtractIntegers` function from the `utils` package.
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter the numbers and keep only the prime numbers using the `Filter` function from the `utils` package.
	result := utils.Filter(numbers, utils.IsPrime)

	// Create a writer to write the output.
	outputWriter := bufio.NewWriter(output)

	// Schedule the flushing of the writer at the end of the function execution.
	defer outputWriter.Flush()

	// Write the sample output to the writer, combining it with the result converted to a string using the `OutputString` function from the `utils` package.
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
