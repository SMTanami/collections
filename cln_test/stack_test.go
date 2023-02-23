package cln_test

import (
	"testing"

	"github.com/SMTanami/collections/cln"
)

func TestStack_Add(t *testing.T) {
	t.Run("Add Should Add Single Element to Stack When Single Element is Given", func(t *testing.T) {
		st := cln.NewStack[int]()
		exp := []int{4}

		st.Add(4)

		valid, msg := ValidateCollection[int](exp, st)
		if !valid {
			t.Errorf("Failed to add value to stack! %s", msg)
		}
	})

	t.Run("Add Should Add Many Elements to Empty Stack When Given Multiple Values", func(t *testing.T) {
		st := cln.NewStack[int]()
		exp := []int{1, 2, 3, 4, 5}

		st.Add(1, 2, 3, 4, 5)

		valid, msg := ValidateCollection[int](exp, st)
		if !valid {
			t.Errorf("Failed to add value to stack! %s", msg)
		}
	})

	t.Run("Add Should Add Many Elements to Non-Empty Stack When Given Multiple Values", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5)
		exp := []int{1, 2, 3, 4, 5, 7}

		st.Add(7)

		valid, msg := ValidateCollection[int](exp, st)
		if !valid {
			t.Errorf("Failed to add value to stack! %s", msg)
		}
	})

	t.Run("Add Should Properly Adjust Size of Stack When Elements are Added", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(4, 5, 6, 7, 8, 9)
		expSize := 6

		actSize := st.Size()
		if actSize != expSize {
			t.Errorf("Add failed to properly resize Stack! Expected %d, got %d", expSize, actSize)
		}
	})
}

func TestStack_Take(t *testing.T) {
	t.Run("Take Should Return Nil when Stack is Empty", func(t *testing.T) {
		st := cln.NewStack[int]()

		val, pop := st.Take()

		if pop != false {
			t.Errorf("Take failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Take Should Return Top of Stack When Stack Contains One Element", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1)

		val, pop := st.Take()

		if pop == false || val != 1 {
			t.Errorf("Take failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Take Should Return Top of Stack When Stack Contains Multiple Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7)
		exp := []int{7, 6, 5, 4, 3, 2, 1}

		act := make([]int, 0, 0)
		for i := 0; i < st.Size(); i++ {
			val, pop := st.Take()
			if pop != true {
				t.Errorf("Got zero value when drawing from top of stack! Expected %d, got %v", exp[i], val)
			}
			act = append(act, val)
		}

		for i, v := range act {
			if v != exp[i] {
				t.Errorf("Failed to return top of stack! Expected %d, got %v", v, exp[i])
			}
		}
	})

	t.Run("Take Should Not Fail When Done Multiple Times ON Empty Stack", func(t *testing.T) {
		st := cln.NewStack[int]()
		expectedOrdering := []int{5, 6, 7, 8}

		st.Add(1)
		for i := 0; i < 5; i++ {
			st.Take()
		}
		st.Add(5, 6, 7, 8)

		isValid, msg := ValidateCollection[int](expectedOrdering, st)
		if !isValid {
			t.Error(msg)
		}
	})

	t.Run("Take Should Properly Decrease Size of Stack When Done Once", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3)
		exp := 2

		st.Take()

		size := st.Size()
		if size != exp {
			t.Errorf("Take failed to resize Stack! Expected %d, got %d", exp, size)
		}
	})

	t.Run("Take Should Properly Decrease Size of Stack When Done Multiple Times", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6)
		exp := 2

		for i := 0; i < 4; i++ {
			st.Take()
		}

		act := st.Size()
		if act != exp {
			t.Errorf("Take failed to resize Stack! Expected %d, got %d", exp, act)
		}
	})
}

func TestStack_Contains(t *testing.T) {
	t.Run("Contains Should Return False When Stack is Empty", func(t *testing.T) {
		st := cln.NewStack[int]()
		val := 7

		if st.Contains(7) {
			t.Errorf("st.Contains(%d) returned true when Stack is empty!", val)
		}
	})

	t.Run("Contains Should Return True When Used On Stack With Single Element", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1)
		val := 1

		if !st.Contains(val) {
			t.Errorf("\nst.Contains(%d) returned false for Stack: %s", val, st.String())
		}
	})

	t.Run("Contains Should Return True When Used On Stack With Multiple Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		desiredValue := 7

		if !st.Contains(desiredValue) {
			t.Errorf("\nst.Contains(%d) returned false for Stack: %s", desiredValue, st.String())
		}
	})
}

func TestStack_Clear(t *testing.T) {
	t.Run("Clear Should Empty the Stack of All Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		expectedOrdering := []int{}
		st.Add(1, 2, 3, 4, 5, 6, 7)

		st.Clear()

		isValid, msg := ValidateCollection[int](expectedOrdering, st)
		if !isValid {
			t.Errorf("Clear did not remove all elements from the Stack! %s", msg)
		}
	})

	t.Run("Clear Should Return Stack with Size of 0", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7)

		st.Clear()

		size := st.Size()
		if size != 0 {
			t.Errorf("Size after clear is %d but expected 0.", size)
		}
	})
}

func TestStack_Filter(t *testing.T) {

	var isGreaterThanTen func(element int) bool = func(stueueElement int) bool {
		return stueueElement > 10
	}

	var isLessThanOneHundred func(element int) bool = func(stueueElement int) bool {
		return stueueElement < 100
	}

	t.Run("Filter Should Not Crash When Stack is Empty", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Filter(isGreaterThanTen)
	})

	t.Run("Stack Should Leave Stack Unchanegd When Filter Is Always False", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		st.Filter(isGreaterThanTen)

		isValid, msg := ValidateCollection[int](exp, st)
		if !isValid {
			t.Errorf("Filter removed elements from stueue when given func is always false! %s", msg)
		}
	})

	t.Run("Stack Should Remove All Elements When Filter Is Always True", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		exp := []int{}

		st.Filter(isLessThanOneHundred)

		isValid, msg := ValidateCollection[int](exp, st)
		if !isValid {
			t.Errorf("Filter did not remove elements from stueue when given func is always true! %s", msg)
		}
	})

	t.Run("RemoevIf Should Reduce the Size of the Stack When it Removes Several Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedSize := 12

		st.Filter(isGreaterThanTen)

		actualSize := st.Size()
		if actualSize != expectedSize {
			t.Errorf("Filter failed to properly resize stueue, expected %d but is %d", expectedSize, actualSize)
		}
	})

	t.Run("Filter Should Maintain Proper Order of Stack Nodes When Filter is Sometimes True", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		exp := []int{2, 7, 3, 8, 1, 9, 1, 3, 5, 5, 5, 1}

		isGreaterThanTen := func(stueueElement int) bool {
			return stueueElement > 10
		}

		st.Filter(isGreaterThanTen)

		isValid, msg := ValidateCollection[int](exp, st)
		if !isValid {
			t.Error(msg)
		}
	})
}

func TestStack_Remove(t *testing.T) {
	t.Run("Remove Should Not Crash When Stack ss Empty", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Remove(4)
	})

	t.Run("Remove Should Leave Stack Unchanged When Stack Does Not Contain Element", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(2, 5, 3, 4, 1, 7)
		expectedOrdering := []int{2, 5, 3, 4, 1, 7}

		st.Remove(6)

		isValid, msg := ValidateCollection[int](expectedOrdering, st)
		if !isValid {
			t.Errorf("Remove removed an element from the stueue when it did not contain the given argument! %s", msg)
		}
	})

	t.Run("Remove Should Remove Element From Stack When Stack Contains Given Element", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(2, 5, 3, 4, 1, 7)
		expectedOrdering := []int{2, 5, 3, 1, 7}

		st.Remove(4)

		isValid, msg := ValidateCollection[int](expectedOrdering, st)
		if !isValid {
			t.Errorf("Remove removed an element from the stueue when it did not contain the given argument! %s", msg)
		}
	})

	t.Run("Remove Should Only Remove First Instance of Given Argument From Stack When Stack Contains Several Instances of Given Element", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(2, 5, 3, 4, 1, 4, 7)
		expectedOrdering := []int{2, 5, 3, 4, 1, 7}

		st.Remove(4)

		isValid, msg := ValidateCollection[int](expectedOrdering, st)
		if !isValid {
			t.Errorf("Remove removed an element from the stueue when it did not contain the given argument! %s", msg)
		}
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("Size Should Return 0 When Stack is Empty", func(t *testing.T) {
		st := cln.NewStack[int]()

		size := st.Size()
		if size != 0 {
			t.Errorf("New Stack has size of %d instead of 0!", size)
		}
	})

	t.Run("Size Should Return 10 When Stack Contains 10 Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		size := st.Size()
		if size != 10 {
			t.Errorf("Stack of 10 elements has size of %d!", size)
		}
	})
}

func TestStack_IsEmpty(t *testing.T) {
	t.Run("IsEmpty Should Return True When Stack is Empty", func(t *testing.T) {
		st := cln.NewStack[int]()

		if !st.IsEmpty() {
			t.Error("New Stack is not empty!")
		}
	})

	t.Run("IsEmpty Should Return False When Stack Contains Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1)

		if st.IsEmpty() {
			t.Error("Stack with single value is returning true for IsEmpty()!")
		}
	})

	t.Run("IsEmpty Should Return False and then True When Stack is Empty and Then an Element is Added", func(t *testing.T) {
		st := cln.NewStack[int]()

		st.Add(1)

		if st.IsEmpty() {
			t.Error("Non-empty Stack is returning true for IsEmpty()!")
		}

		st.Take()

		if !st.IsEmpty() {
			t.Error("Empty Stack's IsEmpty() call returning false!")
		}

		st.Add(1)

		if st.IsEmpty() {
			t.Error("Non-empty Stack is returning true for IsEmpty()!")
		}
	})
}

func TestStack_String(t *testing.T) {
	t.Run("String Should Return Empty String When Stack is Empty", func(t *testing.T) {
		st := cln.NewStack[int]()
		expectedString := "[]"

		actualString := st.String()
		if actualString != expectedString {
			t.Errorf("String did not return empty string when Stack is empty! Expected %s, got %s", expectedString, actualString)
		}
	})

	t.Run("String Return String without Leading Arrow For Last Element", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1)

		expectedString := "[1]"
		actualString := st.String()

		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})

	t.Run("String Should Return Arrow-Linked String When Stack Contains Multiple Elements", func(t *testing.T) {
		st := cln.NewStack[int]()
		st.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedString := "[1 2 3 4 5 6 7 8 9]"
		actualString := st.String()
		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})
}

func TestStack_Type(t *testing.T) {
	t.Run("Queue Should Be a Collection", func(t *testing.T) {
		var c cln.Collection[int]
		c = cln.NewStack[int]()

		_, ok := c.(cln.Collection[int])
		if !ok {
			t.Error("Stack is not a Collection!")
		}
	})
}
