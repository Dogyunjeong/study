package main

import (
	"fmt"

	"14.session28-test-bench--marking/exercise/02/quote"
	"14.session28-test-bench--marking/exercise/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
