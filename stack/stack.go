package stack

import (
	"fmt"
)

type Collection[T comparable] interface {
	Add(elements ...T)
	Draw() (T, bool)
	Contains(element T) bool
	Clear()
	Size() int
	IsEmpty() bool
}

// Filter is the interface that wraps the basic Filter and Remove methods.
type Filterable[T comparable] interface {
	Remove(v T)
	Filter(filter func(v T) bool)
}

// A Stack is a Collection implementation that maintains data in a LIFO (last-in-first-out) manner. All elements added
// to the stack are added to the 'top' of the stack. Operations used to retrieve data return the value stored at the 'top'
// of the stack.
//
// This stack is implemented using a slice, therefore it's size is dynamic.
type stack[T comparable] struct {
	pile []T
}

func New[T comparable]() *stack[T] {
	return &stack[T]{}
}

func (st *stack[T]) Add(elements ...T) {
	st.pile = append(st.pile, elements...)
}

func (st *stack[T]) Draw() (T, bool) {
	if len(st.pile) == 0 {
		var zero T
		return zero, false
	}

	top := st.pile[len(st.pile)-1]
	st.pile = st.pile[:len(st.pile)-1]
	return top, true
}

func (st *stack[T]) Clear() {
	st.pile = make([]T, 0, 0)
}

func (st *stack[T]) Contains(element T) bool {
	for _, v := range st.pile {
		if v == element {
			return true
		}
	}

	return false
}

func (st *stack[T]) Remove(v T) {
	if len(st.pile) > 0 {
		for i := len(st.pile) - 1; i >= 0; i-- {
			if st.pile[i] == v {
				st.pile = append(st.pile[:i], st.pile[i+1:]...)
				return
			}
		}
	}
}

func (st *stack[T]) Filter(filter func(v T) bool) {
	for i := len(st.pile) - 1; i >= 0; i-- {
		if filter(st.pile[i]) {
			st.pile = append(st.pile[:i], st.pile[i+1:]...)
		}
	}
}

func (st *stack[T]) Size() int {
	return len(st.pile)
}

func (st *stack[T]) IsEmpty() bool {
	if len(st.pile) == 0 {
		return true
	}

	return false
}

func (st *stack[T]) String() string {
	return fmt.Sprint(st.pile)
}
