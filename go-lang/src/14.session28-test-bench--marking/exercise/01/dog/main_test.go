package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(7)
	if y != 49 {
		t.Error("Expected", 49, "Got", y)
	}
}
func TestYearsTwo(t *testing.T) {
	y := YearsTwo(7)
	if y != 49 {
		t.Error("Expected", 49, "Got", y)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7 + i)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}

func ExampleYears() {
	y := Years(7)
	fmt.Println(y)
	// Output:
	// 49
}

func ExampleYearsTwo() {
	y := YearsTwo(6)
	fmt.Println(y)
	// Output:
	// 42
}
