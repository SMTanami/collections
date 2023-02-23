package cols

import (
	"fmt"
	"strings"
)

// A Queue is a data structure that (in this implementation) maintains data in a FIFO (first-in-first-out) manner.
// All elements added to the queue are added to the 'tail' end of the queue. Operations used to retrieve data - Take()
// and Peek() - return the value stored at the 'head' of the queue.
//
// This queue is implemented using nodes, not slices or arrays; this decision has it's tradeoffs. A node implementation
// enables the addition and removal of a node to the queue to be O(1) and the memory used by the queue to be O(n) - always.
// On the other hand, a node based implementation is not as performant when adding many values (batches) at a single time consistently.
type Queue[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

// A node is a data object that holds a value and a reference to a following node. Nodes are used
// internally by the queue, and is therefore non-exportable.
type node[T comparable] struct {
	val  T
	next *node[T]
}

// Returns a new instance of a queue of the specified type.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// Adds element(s) to the tail-end of the queue.
func (q *Queue[T]) Add(vals ...T) {
	for _, elem := range vals {
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
func (q *Queue[T]) Take() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	val := q.head.val
	q.head = q.head.next
	q.size--
	return val, true
}

// Returns the value of the head of the queue but does not remove it. If the queue is empty, returns nil.
func (q *Queue[T]) Peek() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	return q.head.val, true
}

// Removes all elements from the queue.
func (q *Queue[T]) Clear() {
	cleared := NewQueue[T]()
	*q = *cleared
}

// Returns true if the queue contains the given element, returns false otherwise.
func (q *Queue[T]) Contains(element T) bool {
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
func (q *Queue[T]) Remove(val T) {
	sentinel := &node[T]{next: q.head}

	for sentinel.next != nil {
		if sentinel.next.val == val {
			sentinel.next = sentinel.next.next
			q.size--
			return
		}

		sentinel = sentinel.next
	}
}

// Filters all elements from the queue that satisfy the given predicate.
func (q *Queue[T]) Filter(filter func(val T) bool) {
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
func (q *Queue[T]) Size() int {
	return q.size
}

// Returns true if the queue contains no elements, otherwise returns false.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Returns a string representation of the queue. The larger the queue, the more expensive the operation.
func (q *Queue[T]) String() string {
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

// Returns a chan of the same type of the collection
func (q *Queue[T]) Iter() chan T {
	c := make(chan T)
	go func() {
		head := q.head
		for head != nil {
			c <- head.val
			head = head.next
		}
		close(c)
	}()
	return c
}
