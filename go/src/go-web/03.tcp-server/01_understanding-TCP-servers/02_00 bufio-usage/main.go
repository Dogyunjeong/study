package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was alwa..."

	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
