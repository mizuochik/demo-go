package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

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

func IdentityInt[V ~int](v V) V {
	return v
}

type MyInt int

type MySet[T comparable] struct {
	d map[T]struct{}
}

func (s *MySet[T]) Add(v T) {
	s.d[v] = struct{}{}
}

func (s *MySet[T]) Has(v T) bool {
	_, ok := s.d[v]
	return ok
}

func main() {
	fmt.Println(Compare(100, 20))
	fmt.Println(Sum([]float64{0, 1, 2, 3, 4, 5}))
	fmt.Println(Add(10.0, 20.0))
	var v int
	Set100(&v)
	fmt.Println(v)
	fmt.Println(IdentityInt(MyInt(100)))
	i, ok := slices.BinarySearch([]int{0, 1, 2, 3, 4, 5}, 2)
	if ok {
		fmt.Printf("found: %d\n", i)
	} else {
		fmt.Println("not found")
	}
	s := &MySet[string]{
		d: make(map[string]struct{}, 0),
	}
	s.Add("hello")
	fmt.Println(s.Has("hello"))
	fmt.Println(s.Has("bye"))
}
