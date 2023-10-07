package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/bouroo/one2n-go-bootcamp/basic-number-filtering/utils"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func main() {
	MainIO(os.Stdin, os.Stdout)
}

// MainIO function takes an input reader and an output writer as parameters.
func MainIO(input io.Reader, output io.Writer) {
	// Prompt the user for input
	fmt.Print("Sample Input: ")

	// Create a scanner to read the input
	inputScanner := bufio.NewScanner(input)
	inputScanner.Scan()

	// Extract the integers from the input
	numbers := utils.ExtractIntegers(inputScanner.Text())

	// Filter the numbers to keep only the even ones
	result := utils.Filter(numbers, utils.IsEven)

	// Create a writer to write the output
	outputWriter := bufio.NewWriter(output)

	// Ensure the writer is flushed at the end
	defer outputWriter.Flush()

	// Write the output string to the writer
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}
