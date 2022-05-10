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

func Add[V interface {
	float32 | float64
}](a, b V) V {
	return a + b
}

func Set100[V int](ref *V) {
	*ref = 100
}

func main() {
	fmt.Println(Compare(100, 20))
	fmt.Println(Sum([]float64{0, 1, 2, 3, 4, 5}))
	fmt.Println(Add(10.0, 20.0))
	var v int
	Set100(&v)
	fmt.Println(v)
}
