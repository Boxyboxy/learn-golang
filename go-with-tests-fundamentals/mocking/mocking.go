package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// sleep will pause execution for the defined Duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	// actual sleeper implementing an interface we have defined above vs a mocked sleeper in the test code
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

/*

Try to make it so your tests are testing useful behaviour unless the implementation is really important to how the system runs.

Some principles/though processes and rule to follow for mocking:
The definition of refactoring is that the code changes but the behaviour stays the same. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes.
So when writing a test ask yourself:
1. Am I testing the behaviour I want, or the implementation details?
2. If I were to refactor this code, would I have to make lots of changes to the tests?

Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour.
Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.

I feel like if a test is working with MORE THAN 3 MOCKS THEN IT IS A RED FLAG - time for a rethink on the design

Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful
but that MEANS A TIGHTER COUPLING BETWEEN TEST CODE AND IMPLEMENTATION coupling between your test code and the implementation.
Be sure you actually care about these details if you're going to spy on them
*/
