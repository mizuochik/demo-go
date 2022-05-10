package main

import "fmt"

func Compare[T comparable](l, r T) bool {
	return l == r
}

type Number interface {
	int64 | float64
}

func Sum[V Number](n []V) V {
	var s V
	for _, nv := range n {
		s += nv
	}
	return s
}

func main() {
	fmt.Println(Compare(100, 20))
	fmt.Println(Sum([]float64{0, 1, 2, 3, 4, 5}))
}