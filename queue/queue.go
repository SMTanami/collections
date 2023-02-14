package queue

import (
	"fmt"
	"strings"
)

// A node is a container that holds a value as well as a reference to the next node after it. Nodes are used internally by the queue, and is therefore non-exportable.
type node[T comparable] struct {
	val  T
	next *node[T]
}

// A Queue is a data structure that (at least in this instance) maintains data in a FIFO (first-in-first-out) manner.
// All elements added to this queue are added to the 'tail' of the queue. Any operations used to retrieve data from
// the queue via Poll() or Peek() returns the 'head' of the queue.
//
// This queue is implemented by using nodes, not slices or arrays. This decision has it's pros and cons. Adding to the
// queue is always O(1) and the memory used by the queue is always O(n) and no more. Queue's that are implemented via
// arrays maintain their empty cells of incredible length, which therefore can lead to an array of 10,000 indexes with
// only a handful of elements. A node based implementation, however, causes slower performance of adding in batches, which
// may be useful in some cases. Another queue implementation will be added for such benefits.
type queue[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Returns a new instance of a queue of the specified type.
func New[T comparable]() *queue[T] {
	return &queue[T]{}
}

// Adds one or many elements to the queue.
func (q *queue[T]) Add(elements ...T) {
	for _, elem := range elements {
		if q.head == nil {
			initialNode := &node[T]{val: elem}
			q.head = initialNode
			q.tail = initialNode
		} else {
			q.tail.next = &node[T]{val: elem, next: nil}
			q.tail = q.tail.next
		}

		q.size++
	}
}

// Returns the value of the head of the queue and removes it. If the queue is empty, returns nil.
func (q *queue[T]) Poll() any {
	if q.IsEmpty() {
		return nil
	}

	val := q.head.val
	q.head = q.head.next
	q.size--
	return val
}

// Returns the value of the head but does not remove it. If the queue is empty, returns nil.
func (q *queue[T]) Peek() any {
	if q.IsEmpty() {
		return nil
	}

	return q.head.val
}

// Removes all of the elements from the queue.
func (q *queue[T]) Clear() {
	cleared := New[T]()
	*q = *cleared
}

// Returns true if the queue contains the given element, returns false otherwise.
func (q *queue[T]) Contains(element T) bool {
	head := q.head
	for head != nil {
		if head.val == element {
			return true
		}
		head = head.next
	}

	return false
}

// Returns the amount of elements contained within the queue.
func (q *queue[T]) Size() int {
	return q.size
}

// Returns true if the queue contains no elements, false otherwise.
func (q *queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Returns a string representation of the queue. The larger the queue, the more expensive the operation.
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
