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

// MainIO function takes an input reader and an output writer as parameters.
func MainIO(input io.Reader, output io.Writer) {

	fmt.Fprint(output, "Sample Input: ")

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	inputText := scanner.Text()

	numbers := utils.ExtractIntegers(inputText)

	fmt.Fprint(output, "Conditions specified using a set of functions: ")
	scanner.Scan()
	conditionsText := scanner.Text()

	conditions := ExtractConditions[int](conditionsText)

	result := utils.FilterAny(numbers, conditions...)

	writer := bufio.NewWriter(output)
	defer writer.Flush()

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
