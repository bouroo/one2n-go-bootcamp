package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

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

	// Prompt the user for conditions
	fmt.Print("Conditions specified using a set of functions: ")
	inputScanner.Scan()

	// Filter the numbers by keeping only the odd ones using a utility function
	result := utils.Filter(numbers, utils.IsOdd)

	// Create a buffered writer to write the output
	outputWriter := bufio.NewWriter(output)

	// Flush the buffered writer to ensure all data is written to the output
	defer outputWriter.Flush()

	// Write the output header and output string
	outputWriter.WriteString("Sample Output: " + utils.OutputString(result) + "\n")
}

func ExtractConditions[T utils.Number](input string) []utils.FilterFn[T] {
	conditions := strings.Split(input, ",")
	result := make([]utils.FilterFn[T], len(conditions))
	for i, condition := range conditions {
		result[i] = ConditionFromString[T](condition)
	}
	return result
}

func ConditionFromString[T utils.Number](input string) utils.FilterFn[T] {
	if strings.Contains(input, "even") {
		return utils.IsEven
	}
	if strings.Contains(input, "odd") {
		return utils.IsOdd
	}
	if strings.Contains(input, "prime") {
		return utils.IsPrime
	}
	reg := regexp.MustCompile(`\d+`)
	options := reg.FindAllString(input, -1)
	if len(options) > 0 {
		optNum, err := strconv.Atoi(options[0])
		if err == nil {
			if strings.Contains(input, "greater") {
				return utils.IsGreaterThan(T(optNum))
			}
			if strings.Contains(input, "less") {
				return utils.IsLessThan(T(optNum))
			}
			if strings.Contains(input, "multiple") {
				return utils.IsMultipleOf(T(optNum))
			}
		}
	}
	return nil
}
