package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dup4()
}

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
// run as an exe or go run ____ then press ctrl+z followed by enter to end the program
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

	// The program then enters another range-based for loop using for line, n := range counts. This loop iterates over each key-value pair in the counts map, where line represents the line (key) and n represents the count (value).
	// Inside this loop, the program checks if the count of a line is greater than 1 using if n > 1. If a line appears more than once, it means it is a duplicate.
	//For each duplicate line, the program prints the count and the line using fmt.Printf("%d\t%s\n", n, line). The count is printed first, followed by a tab character (\t), and then the line itself. The \n at the end of the format string adds a newline character after each printed line.

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts { // range-based for loop
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
func dup2() {
	counts := make(map[string]int)
	// A map is a reference to the data structure created by make. When a map is passed to a function, the function receives a copy of the reference, so any
	// changes the called function makes to the underlying data structure will be visible through the caller’s map reference too.
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

// dup3
func dup3() {

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile returns a byte slice that must be converted int o a string so it can be split by strings.Split
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func dup4() {
	counts2 := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines2(os.Stdin, "os.Stdin", counts2)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines2(f, arg, counts2)
			f.Close()
		}
	}
	for line, filenames := range counts2 {
		fileCount := len(filenames)
		if fileCount == 1 {
			total := 0
			for _, count := range filenames {
				total += count
			}
			if total <= 1 {
				continue
			}
		}

		fmt.Printf("[Found in %d file(s)]\t%s\n", fileCount, line)
		for name, count := range filenames {
			fmt.Printf("\t%d hit(s) in %s\n", count, name)
		}
	}
}

// map{key wordString, value map{key filenameString, value count }}
func countLines2(f *os.File, filename string, counts2 map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts2[input.Text()] == nil {
			counts2[input.Text()] = make(map[string]int)
		}
		counts2[input.Text()][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
