package queue

import (
	"fmt"
	"testing"
)

func validateQueue[T comparable](expectedOrdering []T, q queue[T]) (bool, string) {

	if len(expectedOrdering) != q.size {
		return false, fmt.Sprintf("Queue is invalid due to incorrect sizing! Ordering length = %d, Queue size = %d\nExpected: %v\nGot: %s", len(expectedOrdering), q.Size(), expectedOrdering, q.String())
	}

	var queueVal any
	var expectedVal any
	head := q.head
	for i := 0; i < len(expectedOrdering); i++ {
		expectedVal = expectedOrdering[i]
		queueVal = head.val
		if expectedVal != queueVal {
			return false, fmt.Sprintf("Expected %v but got %v at position %d\nExpected: %v\nGot: %s", expectedVal, queueVal, i, expectedOrdering, q.String())
		}
		head = head.next
	}

	return true, "Valid"
}

func TestValidateQueue(t *testing.T) {
	t.Run("Validate Should Return True When Ordering and Queue Match", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		isValid, msg := validateQueue(s, *q)

		if !isValid {
			t.Fatalf("Expected to validate queue invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return True When Queue and Ordering Are Empty", func(t *testing.T) {
		q := Queue[int]()
		s := []int{}

		isValid, msg := validateQueue(s, *q)

		if !isValid {
			t.Fatalf("Expected to validate queue invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Elements Differ", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s := []int{1, 2, 3, 4, 8, 6, 7, 8, 9, 10}

		isValid, msg := validateQueue(s, *q)

		if isValid {
			t.Fatalf("Expected method to fail but validated queue instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Sizes Differ", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s := []int{1, 2, 3, 4}

		isValid, msg := validateQueue(s, *q)

		if isValid {
			t.Fatalf("Expected method to fail but validated queue instead!\nReturned message: %s", msg)
		}
	})
}

func TestQueue_Add(t *testing.T) {
	t.Run("Add Should Add Single Element to Queue When Single Element is Given", func(t *testing.T) {
		q := Queue[int]()
		expectedOrdering := []int{1}
		q.Add(1)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("Add failed to add single element to queue! %s", msg)
		}
	})

	t.Run("Add Should Add Many Elements to Empty Queue When Given Multiple Values", func(t *testing.T) {
		q := Queue[int]()
		expectedOrdering := []int{1, 2, 3, 4}

		q.Add(1, 2, 3, 4)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Error(msg)
		}
	})

	t.Run("Add Should Add Many Elements to Non-Empty Queue When Given Multiple Values", func(t *testing.T) {
		q := Queue[int]()
		expectedOrdering := []int{14, 62, 33, 1, 23, 27, 52, 6}
		q.Add(14, 62, 33, 1)
		q.Add(23, 27, 52, 6)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Error(msg)
		}
	})

	t.Run("Add Should Properly Adjust Size of Queue When Elements are Added", func(t *testing.T) {
		q := Queue[int]()
		q.Add(4, 5, 6, 7, 8, 9)
		expectedSize := 6

		actualSize := q.Size()
		if actualSize != expectedSize {
			t.Errorf("Add failed to properly resize queue! Expected %d, got %d", expectedSize, actualSize)
		}
	})
}

func TestQueue_Poll(t *testing.T) {
	t.Run("Poll Should Return Nil when Queue is Empty", func(t *testing.T) {
		q := Queue[int]()

		var val any = q.Poll()

		if val != nil {
			t.Fatalf("Poll failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Poll Should Return Head of Queue", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1)
		var val any = q.Poll()

		if val != 1 {
			t.Fatalf("Poll failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Poll Should Maintain Order of Queue When Poll is Done on Single Value Queue", func(t *testing.T) {
		q := Queue[int]()
		expectedOrdering := []int{5, 6, 7, 8}

		q.Add(1)
		q.Poll()
		q.Add(5, 6, 7, 8)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Fatal(msg)
		}
	})

	t.Run("Poll Should Maintain Order of Queue When Done Multiple Times", func(t *testing.T) {
		q := Queue[int]()
		expectedOrdering := []int{5, 6, 7, 8}

		for i := 0; i < 5; i++ {
			q.Poll()
		}
		q.Add(5, 6, 7, 8)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Fatal(msg)
		}
	})

	t.Run("Poll Should Properly Decrease Size of Queue When Done Once", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3)

		q.Poll()

		size := q.Size()
		if size != 2 {
			t.Errorf("Poll failed to resize queue! Expected %d, got %d", 2, size)
		}
	})

	t.Run("Poll Should Properly Decrease Size of Queue When Done Multiple Times", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6)
		expectedSize := 2

		for i := 0; i < 4; i++ {
			q.Poll()
		}

		actualSize := q.Size()
		if actualSize != expectedSize {
			t.Errorf("Poll failed to resize queue! Expected %d, got %d", expectedSize, actualSize)
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("Peek Should Return Nil When Queue is Empty", func(t *testing.T) {
		q := Queue[int]()

		val := q.Peek()
		if val != nil {
			t.Errorf("Peek on empty queue resulted in a non-nil value, %v!", val)
		}
	})

	t.Run("Peek Should Return Head When Queue Contains Single Element", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1)

		val := q.Peek()
		if val != 1 {
			t.Errorf("Peek on queue resulted in an unexpected value, %v!", val)
		}
	})

	t.Run("Peek Should Return Head When Head of Queue is Changed", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3)

		for i := 1; i <= q.size; i++ {
			val := q.Peek()
			if val != i {
				t.Errorf("Peek on queue resulted in an unexpected value!\nExpected: %v\nGot: %v", i, val)
			}
			q.head = q.head.next // Change head of the queue without using q.Poll()
		}
	})
}

func TestQueue_Contains(t *testing.T) {
	t.Run("Contains Should Return False When Queue is Empty", func(t *testing.T) {
		q := Queue[int]()
		val := 7

		if q.Contains(7) {
			t.Errorf("q.Contains(%d) returned true when queue is empty!", val)
		}
	})

	t.Run("Contains Should Return True When Used On Queue With Single Element", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1)
		desiredValue := 1

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}
	})

	t.Run("Contains Should Return True When Used On Queue With Multiple Elements", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		desiredValue := 7

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}
	})
}

func TestQueue_Clear(t *testing.T) {
	t.Run("Clear Should Empty the Queue of All Elements", func(t *testing.T) {
		q := Queue[int]()
		expectedOrdering := []int{}
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("Clear did not remove all elements from the queue! %s", msg)
		}
	})

	t.Run("Clear Should Return Queue with Size of 0", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		size := q.Size()
		if size != 0 {
			t.Errorf("Size after clear is %d but expected 0.", size)
		}
	})
}

func TestQueue_RemoveIf(t *testing.T) {

	var isGreaterThanTen func(element int) bool = func(queueElement int) bool {
		return queueElement > 10
	}

	var isLessThanOneHundred func(element int) bool = func(queueElement int) bool {
		return queueElement < 100
	}

	t.Run("RemoveIf Should Not Crash When Queue is Empty", func(t *testing.T) {
		q := Queue[int]()
		q.RemoveIf(isGreaterThanTen)
	})

	t.Run("Queue Should Leave Queue Unchanegd When Filter Is Always False", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedOrdering := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		q.RemoveIf(isGreaterThanTen)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("RemoveIf removed elements from queue when given func is always false! %s", msg)
		}
	})

	t.Run("Queue Should Remove All Elements When Filter Is Always True", func(t *testing.T) {
		q := Queue[int]()
		q.Add(2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedOrdering := []int{}

		q.RemoveIf(isLessThanOneHundred)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("RemoveIf did not remove elements from queue when given func is always true! %s", msg)
		}
	})

	t.Run("RemoevIf Should Reduce the Size of the Queue When it Removes Several Elements", func(t *testing.T) {
		q := Queue[int]()
		q.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedSize := 12

		q.RemoveIf(isGreaterThanTen)

		actualSize := q.Size()
		if actualSize != expectedSize {
			t.Errorf("RemoveIf failed to properly resize queue, expected %d but is %d", expectedSize, actualSize)
		}
	})

	t.Run("RemoveIf Should Maintain Proper Order of Queue Nodes When Filter is Sometimes True", func(t *testing.T) {
		q := Queue[int]()
		q.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedOrdering := []int{2, 7, 3, 8, 1, 9, 1, 3, 5, 5, 5, 1}

		isGreaterThanTen := func(queueElement int) bool {
			return queueElement > 10
		}

		q.RemoveIf(isGreaterThanTen)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Error(msg)
		}
	})
}

func TestQueue_Remove(t *testing.T) {
	t.Run("Remove Should Not Crash When Queue ss Empty", func(t *testing.T) {
		q := Queue[int]()
		q.Remove(4)
	})

	t.Run("Remove Should Leave Queue Unchanged When Queue Does Not Contain Element", func(t *testing.T) {
		q := Queue[int]()
		q.Add(2, 5, 3, 4, 1, 7)
		expectedOrdering := []int{2, 5, 3, 4, 1, 7}

		q.Remove(6)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("Remove removed an element from the queue when it did not contain the given argument! %s", msg)
		}
	})

	t.Run("Remove Should Remove Element From Queue When Queue Contains Given Element", func(t *testing.T) {
		q := Queue[int]()
		q.Add(2, 5, 3, 4, 1, 7)
		expectedOrdering := []int{2, 5, 3, 1, 7}

		q.Remove(4)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("Remove removed an element from the queue when it did not contain the given argument! %s", msg)
		}
	})

	t.Run("Remove Should Only Remove First Instance of Given Argument From Queue When Queue Contains Several Instances of Given Element", func(t *testing.T) {
		q := Queue[int]()
		q.Add(2, 5, 3, 4, 1, 4, 7)
		expectedOrdering := []int{2, 5, 3, 1, 4, 7}

		q.Remove(4)

		isValid, msg := validateQueue(expectedOrdering, *q)
		if !isValid {
			t.Errorf("Remove removed an element from the queue when it did not contain the given argument! %s", msg)
		}
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("Size Should Return 0 When Queueis Empty", func(t *testing.T) {
		q := Queue[int]()

		size := q.Size()
		if size != 0 {
			t.Errorf("New queue has size of %d instead of 0!", size)
		}
	})

	t.Run("Size Should Return 10 When Queue Contains 10 Elements", func(t *testing.T) {
		q := Queue[int64]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		size := q.Size()
		if size != 10 {
			t.Errorf("Queue of 10 elements has size of %d!", size)
		}
	})
}

func TestQueue_IsEmpty(t *testing.T) {
	t.Run("IsEmpty Should Return True When Queue is Empty", func(t *testing.T) {
		q := Queue[int]()

		if !q.IsEmpty() {
			t.Error("New queue is not empty!")
		}
	})

	t.Run("IsEmpty Should Return False When Queue Contains Elements", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1)

		if q.IsEmpty() {
			t.Error("Queue with single value is returning true for IsEmpty()!")
		}
	})

	t.Run("IsEmpty Should Return False and then True When Queue is Empty and Then an Element is Added", func(t *testing.T) {
		q := Queue[int]()

		q.Add(1)

		if q.IsEmpty() {
			t.Error("Non-empty queue is returning true for IsEmpty()!")
		}

		q.Poll()

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
		q := Queue[int]()
		expectedString := ""

		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("String did not return empty string when queue is empty! Expected %s, got %s", expectedString, actualString)
		}
	})

	t.Run("String Return String without Leading Arrow For Last Element", func(t *testing.T) {
		q := Queue[int]()
		q.Add(1)

		expectedString := "1"
		actualString := q.String()

		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})

	t.Run("String Should Return Arrow-Linked String When Queue Contains Multiple Elements", func(t *testing.T) {
		q := Queue[int64]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedString := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9"
		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})
}
