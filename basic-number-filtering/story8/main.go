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

// MainIO is a function that takes an input reader and an output writer as parameters.
func MainIO(input io.Reader, output io.Writer) {
	// Prompt the user for input
	fmt.Fprint(output, "Sample Input: ")

	// Create a new scanner to read input from the user
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	inputText := scanner.Text()

	// Extract the numbers from the input using a utility function
	numbers := utils.ExtractIntegers(inputText)

	// Prompt the user for conditions
	fmt.Fprint(output, "Conditions specified using a set of functions: ")
	scanner.Scan()
	conditionsText := scanner.Text()

	// Extract the conditions from the input using a utility function
	conditions := ExtractConditions[int](conditionsText)

	// Filter the numbers by keeping only the odd ones using a utility function
	result := utils.FilterAny(numbers, conditions...)

	// Create a buffered writer to write the output
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	// Write the output header and output string
	outputString := "Sample Output: " + utils.OutputString(result) + "\n"
	writer.WriteString(outputString)
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
	reg := utils.RegNumPool.Get().(*regexp.Regexp)
	defer utils.RegNumPool.Put(reg)
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
