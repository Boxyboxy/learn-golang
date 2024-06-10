package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dup1()
}

func dup1() {
	// initialising a map
	counts := make(map[string]int)
	// creates a new Scanner object called input using bufio.NewScanner(os.Stdin). This scanner reads input from the standard input (os.Stdin).
	input := bufio.NewScanner(os.Stdin)

	//The program enters a loop using for input.Scan(). This loop reads input line by line until the end of input is reached or an error occurs.
	//Inside the loop, for each line read by the scanner, the program increments the count of that line in the counts map using counts[input.Text()]++.
	// If a line is encountered for the first time, it is added to the map with a count of 1. If a line is encountered again, its count is incremented.

	for input.Scan() {
		counts[input.Text()]++
	}

	// After the input loop, the program ignores any potential errors from input.Err(), as indicated by the comment // NOTE: ignoring potential errors from input.Err().

	// The program then enters another loop using for line, n := range counts. This loop iterates over each key-value pair in the counts map, where line represents the line (key) and n represents the count (value).
	// Inside this loop, the program checks if the count of a line is greater than 1 using if n > 1. If a line appears more than once, it means it is a duplicate.
	//For each duplicate line, the program prints the count and the line using fmt.Printf("%d\t%s\n", n, line). The count is printed first, followed by a tab character (\t), and then the line itself. The \n at the end of the format string adds a newline character after each printed line.

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// run as an exe or go run ____ then press ctrl+z followed by enter to end the program
