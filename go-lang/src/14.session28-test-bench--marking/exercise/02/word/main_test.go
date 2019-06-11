package word

import (
	"fmt"
	"testing"

	"14.session28-test-bench--marking/exercise/02/quote"
)

func TestCount(t *testing.T) {
	n := Count("one tow three")
	if n != 3 {
		t.Error("Got", n, "Expect", 3)
	}
}
func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
