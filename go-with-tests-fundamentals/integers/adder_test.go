package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	t.Run("2+2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectSum(t, expected, sum)
	})

	t.Run("5+1", func(t *testing.T) {
		sum := Add(5, 1)
		expected := 6

		assertCorrectSum(t, expected, sum)
	})
}

func assertCorrectSum(t testing.TB, expected, sum int) {
	t.Helper() // tells the test suite that this method is a helper method,
	//when the test fails, the line number reported will be in our function call rather than inside our test helper.
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Testable examples will cause the example to appear in the godoc documentation
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
