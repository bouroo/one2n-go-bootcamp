package utils

import (
	"fmt"
	"strings"
)

type NumberStore[T Number] struct {
	Numbers []T
}

func (n NumberStore[T]) String() string {
	var output strings.Builder

	for i, number := range n.Numbers {
		if i > 0 {
			output.WriteString(", ")
		}
		output.WriteString(fmt.Sprint(number))
	}

	return output.String()
}

func (n *NumberStore[T]) FromNumbers(numbers []T) *NumberStore[T] {
	return &NumberStore[T]{
		Numbers: numbers,
	}
}

func (n *NumberStore[T]) EvenFilter() *NumberStore[T] {
	return &NumberStore[T]{
		Numbers: Filter(n.Numbers, IsEven),
	}
}

func (n *NumberStore[T]) OddFilter() *NumberStore[T] {
	return &NumberStore[T]{
		Numbers: Filter(n.Numbers, IsOdd),
	}
}

func (n *NumberStore[T]) PrimeFilter() *NumberStore[T] {
	return &NumberStore[T]{
		Numbers: Filter(n.Numbers, IsPrime),
	}
}

func (n *NumberStore[T]) MultipleOfFilter(multiplyer int) *NumberStore[T] {
	fn := func(number T) bool {
		return int(number)%multiplyer == 0
	}
	return &NumberStore[T]{
		Numbers: Filter(n.Numbers, fn),
	}
}

func (n *NumberStore[T]) GreatherThanFilter(greaterThan int) *NumberStore[T] {
	fn := func(number T) bool {
		return int(number) > greaterThan
	}
	return &NumberStore[T]{
		Numbers: Filter(n.Numbers, fn),
	}
}

func (n *NumberStore[T]) LessThanFilter(lessThan int) *NumberStore[T] {
	fn := func(number T) bool {
		return int(number) < lessThan
	}
	return &NumberStore[T]{
		Numbers: Filter(n.Numbers, fn),
	}
}

func (n *NumberStore[T]) FilterAll(filterFn ...FilterFn[T]) *NumberStore[T] {
	tmpResult := new(NumberStore[T])
numLoop:
	for _, num := range n.Numbers {
		var tmpNum T
		for _, fn := range filterFn {
			if !fn(num) {
				continue numLoop
			}
			tmpNum = num
		}
		tmpResult.Numbers = append(tmpResult.Numbers, tmpNum)
	}
	n = tmpResult
	return n
}
func (n *NumberStore[T]) FilterAny(filterFn ...FilterFn[T]) *NumberStore[T] {
	tmpResult := new(NumberStore[T])
	for _, num := range n.Numbers {
		var tmpNum T
	filterLoop:
		for _, fn := range filterFn {
			if fn(num) {
				tmpNum = num
				break filterLoop
			}
		}
		tmpResult.Numbers = append(tmpResult.Numbers, tmpNum)
	}
	n = tmpResult
	return n
}
