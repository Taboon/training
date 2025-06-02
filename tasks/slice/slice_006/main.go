package main

import "fmt"

func main() {
	a := make([]int, 5, 10)
	b := a[:6]
	fmt.Println(b)
} // 0,0,0,0,0,0
