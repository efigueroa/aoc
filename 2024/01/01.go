package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// slices
	var leftSlice []int
	var rightSlice []int

	// open the file
	file, err := os.Open("../inputs/01")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Store into new slices for left and right numbers
	for scanner.Scan() {
		line := scanner.Text()       // Current line
		bits := strings.Fields(line) // store in slace
		// convert to int
		left, err := strconv.Atoi(bits[0])
		check(err)
		right, err := strconv.Atoi(bits[1])
		check(err)

		// append to slices
		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
	}

	// Sort both in ascending order
	sort.Ints(leftSlice)
	sort.Ints(rightSlice)

	// Loop over both
	sum := 0
	for i := range leftSlice {
		if leftSlice[i] >= rightSlice[i] {
			sum += leftSlice[i] - rightSlice[i]
		} else {
			sum += rightSlice[i] - leftSlice[i]
		}
	}
	fmt.Println(sum)
}
