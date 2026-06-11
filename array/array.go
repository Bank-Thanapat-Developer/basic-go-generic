package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		return *new(T)
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Shift() T {
	if len(*s) == 0 {
		return *new(T)
	}
	first := (*s)[0]
	*s = (*s)[1:]
	return first
}

func (s *Stack[T]) Unshift(item T) {
	*s = append([]T{item}, *s...)
}

func main() {
	arr := Stack[int]{1, 2, 3}

	pop := arr.Pop()
	fmt.Println(pop, arr)

	arr.Push(4)
	fmt.Println(arr)

	shift := arr.Shift()
	fmt.Println(shift, arr)

	arr.Unshift(0)
	fmt.Println(arr)
}
