package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findNPrimes(n int) ([]int, error) {
	if n < 2 {
		return nil, fmt.Errorf("provided number %d is less than 2. No work to be done here", n)
	}

	primes := make([]int, n)
	count := 0
	idx := 2
	for count < n {
		if isPrime(idx) {
			primes[count] = idx
			count++
		}
		idx++
	}

	return primes, nil
}

func isPrime(num int) bool {
	switch {
	case num == 2:
		return true
	case num == 3:
		return true
	case num%2 == 0:
		return false
	case num%3 == 0:
		return false
	default:
		i := 5
		w := 2
		for i*i <= num {
			if num%i == 0 {
				return false
			}
			i += w
			w = 6 - w
		}
		return true
	}
}

func createMatrix(primes []int) [][]int {
	rows := make([][]int, len(primes))
	for i, multiplicand := range primes {
		rows[i] = make([]int, len(primes))
		for j, multiplier := range primes {
			rows[i][j] = multiplicand * multiplier
		}
	}

	return rows
}

func createTable(header []int, matrix [][]int) {
	fmt.Printf("     %-4v\n", header)
	fmt.Println(strings.Repeat("-", (len(header)*5)+5))

	for i, row := range matrix {
		fmt.Printf("%-3d| %-4v\n", header[i], row)
	}
}

func main() {
	var input int
	if len(os.Args) != 2 {
		fmt.Println("Running multiplication table with default N=10.")
		input = 10
	} else {
		var err error
		input, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error parsing parameter. Must be an integer.")
		}
	}
	fmt.Println("Running multiplication table with N=", input)
	result, err := findNPrimes(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	matrix := createMatrix(result)
	createTable(result, matrix)
}
