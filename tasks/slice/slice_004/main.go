package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a)) //cap 3
	copy(b, a)
	b[0] = 42
	fmt.Println("a:", a) //1,2,3
	fmt.Println("b:", b) //42,2,3 при копировании создается новый массив
}
