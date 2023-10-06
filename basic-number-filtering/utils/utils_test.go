package utils_test

import (
	"testing"

	"github.com/bouroo/one2n-go-bootcamp/basic-number-filtering/utils"
	"github.com/stretchr/testify/assert"
)

func TestExtractNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "1, 2, 3, 4, 5",
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    "10,20,30,40,50",
			expected: []int{10, 20, 30, 40, 50},
		},
		{
			input:    "100,200,300,400,500",
			expected: []int{100, 200, 300, 400, 500},
		},
		{
			input:    "1,2,3,abc,5",
			expected: []int{1, 2, 3, 5},
		},
		{
			input:    "1,2,3,4,5,",
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		result := utils.ExtractIntegers(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		number   int
		expected bool
	}{
		{0, true},    // testing 0
		{1, false},   // testing odd number
		{2, true},    // testing even number
		{-1, false},  // testing negative odd number
		{-2, true},   // testing negative even number
		{100, true},  // testing positive even number
		{101, false}, // testing positive odd number
	}

	for _, test := range tests {
		result := utils.IsEven(test.number)
		if result != test.expected {
			t.Errorf("IsEven(%d) = %v, expected %v", test.number, result, test.expected)
		}
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{"Positive odd number", 3, true},
		{"Positive even number", 4, false},
		{"Negative odd number", -5, true},
		{"Negative even number", -6, false},
		{"Zero", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsOdd(tt.number); got != tt.want {
				t.Errorf("IsOdd(%v) = %v, want %v", tt.number, got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		number int
		want   bool
	}{
		{2, true},    // Testing smallest prime number
		{3, true},    // Testing another prime number
		{4, false},   // Testing a non-prime number
		{13, true},   // Testing a larger prime number
		{15, false},  // Testing a non-prime number
		{17, true},   // Testing a larger prime number
		{20, false},  // Testing a non-prime number
		{29, true},   // Testing a larger prime number
		{30, false},  // Testing a non-prime number
		{31, true},   // Testing a larger prime number
		{100, false}, // Testing a non-prime number
	}

	for _, tt := range tests {
		if got := utils.IsPrime(tt.number); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.number, got, tt.want)
		}
	}
}

func TestPrimeFilter(t *testing.T) {
	tests := []struct {
		name          string
		numbers       []int
		expectedPrime []int
	}{
		{
			name:          "Empty List",
			numbers:       []int{},
			expectedPrime: []int(nil),
		},
		{
			name:          "No Prime Numbers",
			numbers:       []int{4, 8, 10},
			expectedPrime: []int(nil),
		},
		{
			name:          "All Prime Numbers",
			numbers:       []int{2, 3, 5, 7, 11},
			expectedPrime: []int{2, 3, 5, 7, 11},
		},
		{
			name:          "Mixed Numbers",
			numbers:       []int{1, 2, 4, 6, 8, 10, 12, 15, 17},
			expectedPrime: []int{2, 17},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualPrime := utils.Filter(test.numbers, utils.IsPrime)
			assert.Equal(t, test.expectedPrime, actualPrime)
		})
	}
}

func TestFilter(t *testing.T) {
	// Test case 1: Filter even numbers
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	filterFn1 := func(num int) bool {
		return num%2 == 0
	}
	expected1 := []int{2, 4, 6}
	result1 := utils.Filter(numbers1, filterFn1)
	if len(result1) != len(expected1) {
		t.Errorf("Test case 1 failed. Expected filtered numbers: %v, but got: %v", expected1, result1)
	}

	// Test case 2: Filter numbers greater than 10
	numbers2 := []int{5, 10, 15, 20, 25}
	filterFn2 := func(num int) bool {
		return num > 10
	}
	expected2 := []int{15, 20, 25}
	result2 := utils.Filter(numbers2, filterFn2)
	if len(result2) != len(expected2) {
		t.Errorf("Test case 2 failed. Expected filtered numbers: %v, but got: %v", expected2, result2)
	}

	// Test case 3: Filter negative numbers
	numbers3 := []int{-3, -2, -1, 0, 1, 2, 3}
	filterFn3 := func(num int) bool {
		return num < 0
	}
	expected3 := []int{-3, -2, -1}
	result3 := utils.Filter(numbers3, filterFn3)
	if len(result3) != len(expected3) {
		t.Errorf("Test case 3 failed. Expected filtered numbers: %v, but got: %v", expected3, result3)
	}
}

func TestOutputString(t *testing.T) {
	testCases := []struct {
		numbers []int
		want    string
	}{
		{[]int{1, 2, 3}, "1, 2, 3"},
		{[]int{4, 5, 6}, "4, 5, 6"},
		{[]int{}, ""},
	}

	for _, tc := range testCases {
		got := utils.OutputString(tc.numbers)
		if got != tc.want {
			t.Errorf("OutputString(%v) = %s, want %s", tc.numbers, got, tc.want)
		}
	}
}
