// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()

	echo2()

	echo3()

	fmt.Println(os.Args[1:])
}

func echo1() {
	var s, sep string // initialises s and sep to empty strings
	// The := sy mbol is part of a short variable declaration,a statement that declares one or more variables and gives them
	// appropriate types based on the initializer values
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// try with running: go run echo.go 1 a b 2
}

func echo2() {
	s, sep := "", ""
	// range loop, each iterationm range produces a pair of values: index and the value of the element at that index
	// _ is used because Go does not permit unused localv ariables, therefore a blank identifier is sued
	for _, arg := range os.Args[1:] {
		// short variable declaration
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/*
	// a traditional "while" loop
	for condition {

	}
*/

/*
	// a traditional infinite loop
	for  {

	}
*/
