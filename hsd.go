// Package foo provide functions for example
package hsd

import "fmt"

// Add returns a plus b.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// Sub returns a minus b.
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		fmt.Printf("a[%d]: %v  b[%d]: %v\n", i, a[i], i, b[i])
		if a[i] != b[i] {
			dist++
		}
	}
	fmt.Printf("distance is %d\n", dist)
	return dist
}
