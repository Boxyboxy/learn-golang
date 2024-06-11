package main

import "fmt"

func Hello() string {
	return "Hello, world!"
}

func HelloName(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello())
}

// go is a compiled language
/*
The Go toolchain conv ertsasource program and the things it
dep ends on int o inst ruc tions in the nat ive machine langu age of a computer. These tools are
accessed through a single command cal le d go that has a number of sub command s. The simplest of these sub command s is run, which compi les the source code fro m on e or more source
files whose names end in .go, lin ks it wit h librar ies, then runs the resulting exe cut able file.
*/
