package cln_test

import (
	"fmt"
	"testing"

	"github.com/SMTanami/collections/cln"
)

func ValidateCollection[T comparable](expOrder []T, st cln.Collection[T]) (bool, string) {

	ss := st.Size()
	if len(expOrder) != ss {
		return false, fmt.Sprintf("Stack is invalid due to incorrect sizing! Ordering length = %d, Stack size = %d\nExpected: %v\nGot: %s", len(expOrder), ss, expOrder, st.String())
	}

	i := 0
	for v := range st.Iter() {
		if expOrder[i] != v {
			return false, fmt.Sprintf("Expected %v but got %v at position %d\nExpected: %v\nGot: %s", expOrder[i], v, i, expOrder, st.String())
		}
		i++
	}

	return true, "Valid"
}

func TestValidateCollection(t *testing.T) {
	t.Run("Validate Should Return True When Ordering and Queue Match", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		valid, msg := ValidateCollection[int](exp, q)

		if !valid {
			t.Errorf("Expected to validate queue invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return True When Queue and Ordering Are Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{}

		valid, msg := ValidateCollection[int](exp, q)

		if !valid {
			t.Errorf("Expected to validate queue invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Elements Differ", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		exp := []int{1, 2, 3, 4, 8, 6, 7, 8, 9, 10}

		valid, msg := ValidateCollection[int](exp, q)

		if valid {
			t.Errorf("Expected method to fail but validated queue instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Sizes Differ", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		exp := []int{1, 2, 3, 4}

		valid, msg := ValidateCollection[int](exp, q)

		if valid {
			t.Errorf("Expected method to fail but validated queue instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return True When Ordering and Stack Match", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		valid, msg := ValidateCollection[int](exp, st)

		if !valid {
			t.Errorf("Expected to validate Stack invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return True When Stack and Ordering Are Empty", func(t *testing.T) {
		st := cln.NewStack[int]()
		exp := []int{}

		isValid, msg := ValidateCollection[int](exp, st)

		if !isValid {
			t.Errorf("Expected to validate Stack invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Elements Differ", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		exp := []int{1, 2, 3, 4, 8, 6, 7, 8, 9, 10}

		isValid, msg := ValidateCollection[int](exp, st)

		if isValid {
			t.Errorf("Expected method to fail but validated Stack instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Sizes Differ", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		exp := []int{1, 2, 3, 4}

		isValid, msg := ValidateCollection[int](exp, st)

		if isValid {
			t.Errorf("Expected method to fail but validated Stack instead!\nReturned message: %s", msg)
		}
	})
}
