package main

import "fmt"

// You can group constants in a block in go.
const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world!"
	}

	return greetingPrefix(language) + name

}

// in Go, public methods start with capital letters
// private methods start with lower letters
func greetingPrefix(language string) (prefix string) { //named return value in function signature
	// variable is created in function and can be returned in function just by calling return.
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}

// go is a compiled language
/*
The Go toolchain conv ertsasource program and the things it
dep ends on int o inst ruc tions in the nat ive machine langu age of a computer. These tools are
accessed through a single command cal le d go that has a number of sub command s. The simplest of these sub command s is run, which compi les the source code fro m on e or more source
files whose names end in .go, lin ks it wit h librar ies, then runs the resulting exe cut able file.
*/
