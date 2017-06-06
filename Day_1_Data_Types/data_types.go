package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var secondInteger uint64
	var secondFloat float64
	var secondString string

	// Read and save an integer, double, and String to your variables.
	if scanner.Scan() {
		secondInteger, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	}

	if scanner.Scan() {
		secondFloat, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	if scanner.Scan() {
		secondString = scanner.Text()
	}
	// Print the sum of both integer variables on a new line.
	fmt.Println(i + secondInteger)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+secondFloat)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + secondString)
}
