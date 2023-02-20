package queue

import (
	"fmt"
	"strings"
)

type Collection[Type comparable] interface {
	Add(elements ...Type)
	Remove(element Type)
	Contains(element Type) bool
	Clear()
	Size() int
	IsEmpty() bool
}

// Filter is the interface that wraps the basic Filter method.
type Filter[Type comparable] interface {
	Filter(filter func(element Type) bool)
}

// A Queue is a data structure that (in this implementation) maintains data in a FIFO (first-in-first-ouType) manner.
// All elements added to the queue are added to the 'tail' end of the queue. Operations used to retrieve data - Poll() and Peek() -
// return the value stored in the 'head' of the queue.
//
// This queue is implemented using nodes, not slices or arrays; this decision has it's pros and cons. Adding to the
// queue is always O(1) and the memory used by the queue is always O(n). Queue's that are implemented using
// arrays or slices maintain many empty array cells when the head is relocated and the pointer to the head is referencing the 10,000th index,
// thus allowing an array or slice of 100,000 indexes to hold just 10,000 elements. On the other hand, a node based implementation
// is not as performant when adding many values at a single time (batches) consistently.
//
// Therefore, another queue implementation will be added to cater to such a use case.
type queue[Type comparable] struct {
	head *node[Type]
	tail *node[Type]
	size int
}

// A node is a data object that holds a value and a reference to a following node. Nodes are used
// internally by the queue, and is therefore non-exportable.
type node[Type comparable] struct {
	val  Type
	next *node[Type]
}

// Returns a new instance of a queue of the specified type.
func New[Type comparable]() *queue[Type] {
	return &queue[Type]{}
}

// Adds element(s) to the tail-end of the queue.
func (q *queue[Type]) Add(elements ...Type) {
	for _, elem := range elements {
		if q.head != nil {
			q.tail.next = &node[Type]{val: elem, next: nil}
			q.tail = q.tail.next
		} else {
			initialNode := &node[Type]{val: elem}
			q.head = initialNode
			q.tail = initialNode
		}

		q.size++
	}
}

// Returns the value of the head of the queue and removes it. If the queue is empty, returns nil.
func (q *queue[Type]) Poll() any {
	if q.IsEmpty() {
		return nil
	}

	val := q.head.val
	q.head = q.head.next
	q.size--
	return val
}

// Returns the value of the head of the queue but does not remove it. If the queue is empty, returns nil.
func (q *queue[Type]) Peek() any {
	if q.IsEmpty() {
		return nil
	}

	return q.head.val
}

// Removes all elements from the queue.
func (q *queue[Type]) Clear() {
	cleared := New[Type]()
	*q = *cleared
}

// Returns true if the queue contains the given element, returns false otherwise.
func (q *queue[Type]) Contains(element Type) bool {
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
func (q *queue[Type]) Remove(element Type) {
	sentinel := &node[Type]{next: q.head}

	for sentinel.next != nil {
		if sentinel.next.val == element {
			sentinel.next = sentinel.next.next
			q.size--
			return
		}

		sentinel = sentinel.next
	}
}

// Filters all elements from the queue that satisfy the given predicate.
func (q *queue[Type]) Filter(filter func(queueElement Type) bool) {
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
func (q *queue[Type]) Size() int {
	return q.size
}

// Returns true if the queue contains no elements, otherwise returns false.
func (q *queue[Type]) IsEmpty() bool {
	return q.size == 0
}

// Returns a string representation of the queue. The larger the queue, the more expensive the operation.
func (q *queue[Type]) String() string {
	var stringBuilder strings.Builder
	head := q.head

	for head != nil {
		if head.next != nil {
			stringBuilder.WriteString(fmt.Sprintf("%v -> ", head.val))
		} else {
			stringBuilder.WriteString(fmt.Sprint(head.val))
		}
		head = head.next
	}

	return stringBuilder.String()
}
