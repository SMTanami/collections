package queue

import (
	"fmt"
	"strings"
)

// A node is a data object that holds a value and a reference to another node that follows it. Nodes are used
// internally by the queue, and is therefore non-exportable.
type node[T comparable] struct {
	val  T
	next *node[T]
}

// A Queue is a data structure that (in this implementation) maintains data in a FIFO (first-in-first-out) manner.
// All elements added to the queue are added to the 'tail' of the queue. Operations used to retrieve data
// such as Poll() or Peek() return the value stored in the 'head' of the queue.
//
// This queue is implemented using nodes, not slices or arrays. This decision has it's pros and cons. Adding to the
// queue is always O(1) and the memory used by the queue is always O(n). Queue's that are implemented using
// arrays or slices maintain many empty cells when the head is relocated, allowing an array or slice of 100,000
// indexes to hold just 1,000 elements. On the other hand, a node based implementation is not as performant when
// adding many values at once consistently. Another type of queue implementation will be added for that use case.
type queue[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Returns a new instance of a queue of the specified type.
func New[T comparable]() *queue[T] {
	return &queue[T]{}
}

// Adds element(s) to the tail-end of the queue.
func (q *queue[T]) Add(elements ...T) {
	for _, elem := range elements {
		if q.head != nil {
			q.tail.next = &node[T]{val: elem, next: nil}
			q.tail = q.tail.next
		} else {
			initialNode := &node[T]{val: elem}
			q.head = initialNode
			q.tail = initialNode
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

// Returns the value of the head of the queue but does not remove it. If the queue is empty, returns nil.
func (q *queue[T]) Peek() any {
	if q.IsEmpty() {
		return nil
	}

	return q.head.val
}

// Removes all elements from the queue.
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

// Removes the first instance of the given element from the queue.
func (q *queue[T]) Remove(element T) {
	sentinel := &node[T]{next: q.head}

	for sentinel.next != nil {
		if sentinel.next.val == element {
			sentinel.next = sentinel.next.next
			q.size--
			return
		}

		sentinel = sentinel.next
	}
}

// Removes all elements that cause the given predicate to output 'true' when used as input.
func (q *queue[T]) RemoveIf(filter func(queueElement T) bool) {
	if q.head == nil {
		return
	}

	currNode := q.head
	for currNode.next != nil {
		if filter(currNode.next.val) {
			currNode.next = currNode.next.next
			q.size--
		} else {
			currNode = currNode.next
		}
	}

	if filter(q.head.val) {
		q.head = q.head.next
		q.size--
	}
}

// Returns the amount of elements contained within the queue.
func (q *queue[T]) Size() int {
	return q.size
}

// Returns true if the queue contains no elements, returns false otherwise.
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
