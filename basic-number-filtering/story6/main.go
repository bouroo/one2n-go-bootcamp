package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bouroo/one2n-go-bootcamp/basic-number-filtering/utils"
)

func main() {
	fmt.Print("Sample Input: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	numbers := utils.ExtractNumbers(input.Text())
	result := utils.Filter(numbers, utils.IsOdd)
	result = utils.Filter(result, IsMultipleByThree)
	result = utils.Filter(result, IsGreatherThanTen)

	fmt.Printf("Sample Output: %s\n", utils.OutputString(result))
}

func IsMultipleByThree(number int) bool {
	return number%3 == 0
}

func IsGreatherThanTen(number int) bool {
	return number > 10
}
