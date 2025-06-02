package main

import "fmt"

func appendSlice(s []int, v int) {
	s = append(s, v)
}

func main() {
	x := make([]int, 1, 3)
	fmt.Println(x)
	appendSlice(x, 1)
	fmt.Println(x)
}
