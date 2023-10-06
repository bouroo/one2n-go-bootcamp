package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainIO(t *testing.T) {
	// Test Case 1: Standard Input and Output
	stdin := bytes.NewBufferString("1, 2, 3, 4, 5, 6, 7, 8, 9, 10")
	stdout := &bytes.Buffer{}
	MainIO(stdin, stdout)
	expected := "Sample Output: 3, 5, 7" + "\n"
	assert.Equal(t, expected, stdout.String())

	// Test Case 2: Empty Input and Empty Output
	stdin = bytes.NewBufferString("")
	stdout = &bytes.Buffer{}
	MainIO(stdin, stdout)
	expected = "Sample Output: " + "\n"
	assert.Equal(t, expected, stdout.String())
}
