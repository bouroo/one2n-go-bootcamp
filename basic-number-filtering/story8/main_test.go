package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainIO(t *testing.T) {
	// Test Case 1: Standard Input and Output
	stdin := bytes.NewBufferString("1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20\n")
	stdin.WriteString("prime, greater than 15, multiple of 5")
	stdout := &bytes.Buffer{}
	MainIO(stdin, stdout)
	expected := "Sample Input: Conditions specified using a set of functions: Sample Output: 2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20" + "\n"
	assert.Equal(t, expected, stdout.String())

	// Test Case 2: Standard Input and Output
	stdin = bytes.NewBufferString("1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20\n")
	stdin.WriteString("less than 6, multiple of 3")
	stdout = &bytes.Buffer{}
	MainIO(stdin, stdout)
	expected = "Sample Input: Conditions specified using a set of functions: Sample Output: 1, 2, 3, 4, 5, 6, 9, 12, 15, 18" + "\n"
	assert.Equal(t, expected, stdout.String())

	// Test Case 3: Empty Input and Empty Output
	stdin = bytes.NewBufferString("\n")
	stdout = &bytes.Buffer{}
	MainIO(stdin, stdout)
	expected = "Sample Input: Conditions specified using a set of functions: Sample Output: " + "\n"
	assert.Equal(t, expected, stdout.String())
}
