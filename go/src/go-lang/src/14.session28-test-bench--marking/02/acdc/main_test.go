package acdc

import "fmt"

// Go test will understand comment of Output and generated, Once it is 7
/**
--- FAIL: ExampleSum (0.00s)
got:
5
want:
7
FAIL
*/
func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}
