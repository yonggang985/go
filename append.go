package main

import "fmt"

func pirntSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func main() {
	var a []int
	pirntSlice("a", a)

	a = append(a, 0)
	pirntSlice("a", a)

	a = append(a, 1)
	pirntSlice("a", a)

	a = append(a, 2, 3, 4)
	pirntSlice("a", a)

}
