package queue

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	val  T
	next *node[T]
}

type queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Add(items ...T) error {
	for _, item := range items {
		if q.head == nil {
			initialNode := &node[T]{val: item}
			q.head = initialNode
			q.tail = initialNode
		} else {
			q.tail.next = &node[T]{val: item, next: nil}
			q.tail = q.tail.next
		}

		q.size++
	}

	return nil
}

func (q *queue[T]) Peek() any {
	if q.head == nil {
		return nil
	}

	return q.head.val
}

func (q *queue[T]) Poll() any {
	if q.size == 0 {
		return nil
	}

	val := q.head.val
	q.head = q.head.next
	q.size--
	return val
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *queue[T]) String() string {
	var stringBuilder strings.Builder
	head := q.head

	for head != nil {
		if head.next == nil {
			stringBuilder.WriteString(fmt.Sprint(head.val))
		} else {
			stringBuilder.WriteString(fmt.Sprintf("%v -> ", head.val))
		}
		head = head.next
	}

	return stringBuilder.String()
}
