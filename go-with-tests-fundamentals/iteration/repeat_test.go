package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// go test -bench=.
// .. When the benchmark code is executed, it runs b.N times and measures how long it takes
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A")
	}
}

// This output comment block actually affects the example
func ExampleRepeat() {
	result := Repeat("a")
	fmt.Println(result)
	// Output: aaaaa
}
