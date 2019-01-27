package dog

import "testing"

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

fuc BenchmarkYears(t *testing.B) {
	
}