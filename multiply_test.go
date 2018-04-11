package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var primeTestCases = []struct {
	input    int
	expected bool
}{
	{2, true},
	{3, true},
	{5, true},
	{7, true},
	{10, false},
	{21, false},
	{113, true},
	{114, false},
	{115, false},
}

var testFindNPrimes = []struct {
	input    int
	expected []int
	errors   bool
}{
	{
		input:    5,
		expected: []int{2, 3, 5, 7, 11},
		errors:   false,
	},
	{
		input:    10,
		expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		errors:   false,
	},
	{
		input:    16,
		expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53},
		errors:   false,
	},
	{
		input:    -2,
		expected: nil,
		errors:   true,
	},
}

var testCreateMatrix = []struct {
	input    []int
	expected [][]int
}{
	{
		input:    []int{2, 4},
		expected: [][]int{{4, 8}, {8, 16}},
	},
}

func TestIsPrime(t *testing.T) {
	for _, tc := range primeTestCases {
		actual := isPrime(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestFindNPrimes(t *testing.T) {
	for _, tc := range testFindNPrimes {
		actual, err := findNPrimes(tc.input)
		assert.Equal(t, tc.expected, actual)
		if tc.errors == true {
			assert.Error(t, err)
		}
	}
}

func TestCreateMatrix(t *testing.T) {
	for _, tc := range testCreateMatrix {
		actual := createMatrix(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}
