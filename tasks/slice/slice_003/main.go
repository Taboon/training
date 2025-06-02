package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5} // len 5 cap 8
	y := x[:2]                // 1, 2, 3 len 2 cap 8
	y = append(y, 99)         // 1, 2, 99
	fmt.Println("x:", x)      // 1, 2, 99, 4, 5
	fmt.Println("y:", y)      // 1, 2, 99
}
