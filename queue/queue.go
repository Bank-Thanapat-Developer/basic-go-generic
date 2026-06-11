package queue

import "fmt"

type Queue[T any] []T

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(*q) == 0 {
		return *new(T)
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func main() {
	queue := Queue[int]{1, 2, 3}
	queue.Enqueue(4)
	fmt.Println(queue)

	dequeue := queue.Dequeue()
	fmt.Println(dequeue, queue)
}
