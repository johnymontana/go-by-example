package main

import (
	"fmt"
	"slices"
)

func main() {

	// typed only by elements, not length
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// make --> init slice
	// cap --> capacity
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// append returns a new slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println("len", len(s))

	c := make([]string, len(s))
	// copy into c from s
	copy(c, s)
	fmt.Println("cpy:", c)

	// get elements 2, 3, 4 (excludes 5)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// up to but excludes 5
	l = s[:5]
	fmt.Println("sl2:", l)

	// up from (and including 2)
	l = s[2:]
	fmt.Println("sl3", l)

	// if this had number of elements, would it be an array?
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// equality based on items?
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d ", twoD)

}
