package cln_test

import (
	"testing"

	"github.com/SMTanami/collections/cln"
)

func TestQueue_Add(t *testing.T) {
	t.Run("Add Should Add Single Element to Queue When Single Element is Given", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{1}
		q.Add(1)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Add failed to add single element to queue! %s", msg)
		}
	})

	t.Run("Add Should Add Many Elements to Empty Queue When Given Multiple Values", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{1, 2, 3, 4}

		q.Add(1, 2, 3, 4)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Add Should Add Many Elements to Non-Empty Queue When Given Multiple Values", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{14, 62, 33, 1, 23, 27, 52, 6}
		q.Add(14, 62, 33, 1)
		q.Add(23, 27, 52, 6)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Add Should Properly Adjust Size of Queue When Elements are Added", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(4, 5, 6, 7, 8, 9)
		expectedSize := 6

		actualSize := q.Size()
		if actualSize != expectedSize {
			t.Errorf("Add failed to properly resize queue! Expected %d, got %d", expectedSize, actualSize)
		}
	})
}

func TestQueue_Take(t *testing.T) {
	t.Run("Take Should Return Nil when Queue is Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()

		val, drew := q.Take()

		if drew != false && val != 0 {
			t.Errorf("Take failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Take Should Return Head of Queue", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1)

		val, drew := q.Take()

		if drew != true || val != 1 {
			t.Errorf("Take failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Take Should Maintain Order of Queue When Take is Done on Single Value Queue", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{5, 6, 7, 8}

		q.Add(1)
		q.Take()
		q.Add(5, 6, 7, 8)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Take Should Maintain Order of Queue When Done Multiple Times", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{5, 6, 7, 8}

		for i := 0; i < 5; i++ {
			q.Take()
		}
		q.Add(5, 6, 7, 8)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Take Should Properly Decrease Size of Queue When Done Once", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3)

		q.Take()

		size := q.Size()
		if size != 2 {
			t.Errorf("Take failed to resize queue! Expected %d, got %d", 2, size)
		}
	})

	t.Run("Take Should Properly Decrease Size of Queue When Done Multiple Times", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6)
		expectedSize := 2

		for i := 0; i < 4; i++ {
			q.Take()
		}

		actualSize := q.Size()
		if actualSize != expectedSize {
			t.Errorf("Take failed to resize queue! Expected %d, got %d", expectedSize, actualSize)
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("Peek Should Return Nil When Queue is Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()

		val, peek := q.Peek()
		if peek != false && val != 0 {
			t.Errorf("Peek on empty queue resulted in a non-nil value, %v!", val)
		}
	})

	t.Run("Peek Should Return Head When Queue Contains Single Element", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1)

		val, peek := q.Peek()
		if peek != true || val != 1 {
			t.Errorf("Peek on queue resulted in an unexpected value, %v!", val)
		}
	})

	t.Run("Peek Should Return Head When Head of Queue is Changed", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3)

		for i := 1; i <= q.Size(); i++ {
			val, peek := q.Peek()
			if peek != true || val != i {
				t.Errorf("Peek on queue resulted in an unexpected value!\nExpected: %v\nGot: %v", i, val)
			}
			q.Take()
		}
	})
}

func TestQueue_Contains(t *testing.T) {
	t.Run("Contains Should Return False When Queue is Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()
		val := 7

		if q.Contains(7) {
			t.Errorf("q.Contains(%d) returned true when queue is empty!", val)
		}
	})

	t.Run("Contains Should Return True When Used On Queue With Single Element", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1)
		desiredValue := 1

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}
	})

	t.Run("Contains Should Return True When Used On Queue With Multiple Elements", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		desiredValue := 7

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}
	})
}

func TestQueue_Clear(t *testing.T) {
	t.Run("Clear Should Empty the Queue of All Elements", func(t *testing.T) {
		q := cln.NewQueue[int]()
		exp := []int{}
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Clear did not remove all elements from the queue! %s", msg)
		}
	})

	t.Run("Clear Should Return Queue with Size of 0", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		size := q.Size()
		if size != 0 {
			t.Errorf("Size after clear is %d but expected 0.", size)
		}
	})
}

func TestQueue_Filter(t *testing.T) {

	var isGreaterThanTen func(element int) bool = func(queueElement int) bool {
		return queueElement > 10
	}

	var isLessThanOneHundred func(element int) bool = func(queueElement int) bool {
		return queueElement < 100
	}

	t.Run("Filter Should Not Crash When Queue is Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Filter(isGreaterThanTen)
	})

	t.Run("Queue Should Leave Queue Unchanegd When Filter Is Always False", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		q.Filter(isGreaterThanTen)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Filter removed elements from queue when given func is always false! %s", msg)
		}
	})

	t.Run("Queue Should Remove All Elements When Filter Is Always True", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		exp := []int{}

		q.Filter(isLessThanOneHundred)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Filter did not remove elements from queue when given func is always true! %s", msg)
		}
	})

	t.Run("RemoevIf Should Reduce the Size of the Queue When it Removes Several Elements", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedSize := 12

		q.Filter(isGreaterThanTen)

		actualSize := q.Size()
		if actualSize != expectedSize {
			t.Errorf("Filter failed to properly resize queue, expected %d but is %d", expectedSize, actualSize)
		}
	})

	t.Run("Filter Should Maintain Proper Order of Queue Nodes When Filter is Sometimes True", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		exp := []int{2, 7, 3, 8, 1, 9, 1, 3, 5, 5, 5, 1}

		isGreaterThanTen := func(queueElement int) bool {
			return queueElement > 10
		}

		q.Filter(isGreaterThanTen)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Error(msg)
		}
	})
}

func TestQueue_Remove(t *testing.T) {
	t.Run("Remove Should Not Crash When Queue ss Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Remove(4)
	})

	t.Run("Remove Should Leave Queue Unchanged When Queue Does Not Contain Element", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(2, 5, 3, 4, 1, 7)
		exp := []int{2, 5, 3, 4, 1, 7}

		q.Remove(6)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Remove removed an element from the queue when it did not contain the given argument! %s", msg)
		}
	})

	t.Run("Remove Should Remove Element From Queue When Queue Contains Given Element", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(2, 5, 3, 4, 1, 7)
		exp := []int{2, 5, 3, 1, 7}

		q.Remove(4)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Remove removed an element from the queue when it did not contain the given argument! %s", msg)
		}
	})

	t.Run("Remove Should Only Remove First Instance of Given Argument From Queue When Queue Contains Several Instances of Given Element", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(2, 5, 3, 4, 1, 4, 7)
		exp := []int{2, 5, 3, 1, 4, 7}

		q.Remove(4)

		valid, msg := ValidateCollection[int](exp, q)
		if !valid {
			t.Errorf("Remove removed an element from the queue when it did not contain the given argument! %s", msg)
		}
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("Size Should Return 0 When Queueis Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()

		size := q.Size()
		if size != 0 {
			t.Errorf("New queue has size of %d instead of 0!", size)
		}
	})

	t.Run("Size Should Return 10 When Queue Contains 10 Elements", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		size := q.Size()
		if size != 10 {
			t.Errorf("Queue of 10 elements has size of %d!", size)
		}
	})
}

func TestQueue_IsEmpty(t *testing.T) {
	t.Run("IsEmpty Should Return True When Queue is Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()

		if !q.IsEmpty() {
			t.Error("New queue is not empty!")
		}
	})

	t.Run("IsEmpty Should Return False When Queue Contains Elements", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1)

		if q.IsEmpty() {
			t.Error("Queue with single value is returning true for IsEmpty()!")
		}
	})

	t.Run("IsEmpty Should Return False and then True When Queue is Empty and Then an Element is Added", func(t *testing.T) {
		q := cln.NewQueue[int]()

		q.Add(1)

		if q.IsEmpty() {
			t.Error("Non-empty queue is returning true for IsEmpty()!")
		}

		q.Take()

		if !q.IsEmpty() {
			t.Error("Empty Queue's IsEmpty() call returning false!")
		}

		q.Add(1)

		if q.IsEmpty() {
			t.Error("Non-empty queue is returning true for IsEmpty()!")
		}
	})
}

func TestQueue_String(t *testing.T) {
	t.Run("String Should Return Empty String When Queue is Empty", func(t *testing.T) {
		q := cln.NewQueue[int]()
		expectedString := ""

		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("String did not return empty string when queue is empty! Expected %s, got %s", expectedString, actualString)
		}
	})

	t.Run("String Return String without Leading Arrow For Last Element", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1)

		expectedString := "1"
		actualString := q.String()

		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})

	t.Run("String Should Return Arrow-Linked String When Queue Contains Multiple Elements", func(t *testing.T) {
		q := cln.NewQueue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedString := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9"
		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})
}

func TestQueue_Type(t *testing.T) {
	t.Run("Queue Should Be a Collection", func(t *testing.T) {
		var c cln.Collection[int]
		c = cln.NewQueue[int]()

		_, ok := c.(cln.Collection[int])
		if !ok {
			t.Error("Queue is not a Collection!")
		}
	})
}
