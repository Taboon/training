package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 100
}

func main() {
	x := []int{1, 2, 3}
	modifySlice(x)
	fmt.Println(x) // Что будет выведено? 100, 2, 3
}
