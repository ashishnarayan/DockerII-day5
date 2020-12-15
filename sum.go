package main

import (
	"fmt"
)

func main() {
	fmt.Printf("sum(100) = %d\n", sum(100))
}

func gaussSum(n int) int {
	fmt.Printf("Doing gaussSum for %d", n)
	return (n * (n + 1)) / 2
}

func sum(n int) int {
	return gaussSum(n)
}
