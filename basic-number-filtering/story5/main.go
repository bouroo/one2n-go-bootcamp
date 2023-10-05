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
	result := utils.Filter(numbers, utils.IsEven)
	result = utils.Filter(result, IsMultipleByFive)

	fmt.Printf("Sample Output: %s\n", utils.OutputString(result))
}

func IsMultipleByFive(number int) bool {
	return number%5 == 0
}
