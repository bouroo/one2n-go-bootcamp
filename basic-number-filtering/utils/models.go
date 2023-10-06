package utils

type NumberStore[T Number] struct {
	Numbers []T
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
