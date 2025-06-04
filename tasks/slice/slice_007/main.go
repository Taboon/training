package main

import "fmt"

func extend(s []int) []int {
	return append(s, 42)
}

func main() {
	x := []int{1, 2}
	y := extend(x)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
