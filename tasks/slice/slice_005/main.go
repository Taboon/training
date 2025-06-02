package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
	} // 1,1 2,2 3,4 4,4 5,8 6,8 7,8 8,8 9,16 10,16
}
