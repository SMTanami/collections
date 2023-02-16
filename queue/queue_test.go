package queue

import (
	"fmt"
	"testing"
)

func validateQueue[T comparable](expectedOrdering []T, q queue[T]) (bool, string) {

	if len(expectedOrdering) != q.Size() {
		return false, fmt.Sprintf("Given ordering and queue are not the same size! Ordering length = %d, Queue size = %d\nExpected: %v\nGot: %s", len(expectedOrdering), q.Size(), expectedOrdering, q.String())
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
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		valid, msg := validateQueue(s, *q)

		if !valid {
			t.Fatalf("Expected to validate queue invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return True When Queue and Ordering Are Empty", func(t *testing.T) {
		q := New[int]()
		s := []int{}

		valid, msg := validateQueue(s, *q)

		if !valid {
			t.Fatalf("Expected to validate queue invalidated it instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Elements Differ", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s := []int{1, 2, 3, 4, 8, 6, 7, 8, 9, 10}

		valid, msg := validateQueue(s, *q)

		if valid {
			t.Fatalf("Expected method to fail but validated queue instead!\nReturned message: %s", msg)
		}
	})

	t.Run("Validate Should Return False When Sizes Differ", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s := []int{1, 2, 3, 4}

		valid, msg := validateQueue(s, *q)

		if valid {
			t.Fatalf("Expected method to fail but validated queue instead!\nReturned message: %s", msg)
		}
	})
}

func TestQueue_Add(t *testing.T) {
	t.Run("Add One to Empty Queue", func(t *testing.T) {
		q := New[int64]()
		q.Add(1)
		if q.IsEmpty() {
			t.Error("Failed to add element to queue!")
		}
	})

	t.Run("Add Many to Empty Queue", func(t *testing.T) {
		q := New[int]()
		expectedOrdering := []int{1, 2, 3, 4}
		q.Add(1, 2, 3, 4)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Add Many to Queue with Elements", func(t *testing.T) {
		q := New[int]()
		expectedOrdering := []int{14, 62, 33, 1, 23, 27, 52, 6}
		q.Add(14, 62, 33, 1)
		q.Add(23, 27, 52, 6)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Error(msg)
		}
	})
}

func TestQueue_Poll(t *testing.T) {
	t.Run("Poll Should Return Head of Queue", func(t *testing.T) {
		q := New[int]()
		q.Add(1)
		var val any = q.Poll()

		if val != 1 {
			t.Fatalf("Poll failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Poll Should Maintain Order of Queue", func(t *testing.T) {
		q := New[int]()
		expectedOrdering := []int{5, 6, 7, 8}

		q.Add(1)
		q.Poll()
		q.Add(5, 6, 7, 8)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Fatal(msg)
		}
	})

	t.Run("Multiple Polls Should Maintain Order of Queue", func(t *testing.T) {
		q := New[int]()
		expectedOrdering := []int{5, 6, 7, 8}

		for i := 0; i < 5; i++ {
			q.Poll()
		}
		q.Add(5, 6, 7, 8)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Fatal(msg)
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("Peek on Empty Queue", func(t *testing.T) {
		q := New[int]()

		val := q.Peek()
		if val != nil {
			t.Errorf("Peek on empty queue resulted in a non-nil value, %v!", val)
		}
	})

	t.Run("Peek on Queue with Single Value", func(t *testing.T) {
		q := New[int]()
		q.Add(1)

		val := q.Peek()
		if val != 1 {
			t.Errorf("Peek on queue resulted in an unexpected value, %v!", val)
		}
	})

	t.Run("Peek on Queue with Multiple Values", func(t *testing.T) {
		q := New[int64]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)

		for i := 1; i <= q.size; i++ {
			val := q.Peek()
			if val != int64(i) {
				t.Errorf("Peek on queue resulted in an unexpected value!\nExpected: %v\nGot: %v", i, val)
			}
			q.Poll()
		}
	})
}

func TestQueue_Contains(t *testing.T) {
	t.Run("Contains on Empty Queue", func(t *testing.T) {
		q := New[int64]()
		val := 7

		if q.Contains(7) {
			t.Errorf("q.Contains(%d) returned true when queue is empty!", val)
		}
	})

	t.Run("Contains on Single Element Queue", func(t *testing.T) {
		q := New[int]()
		q.Add(1)
		desiredValue := 1

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}
	})

	t.Run("Contains on Multi-Element Queue", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		desiredValue := 7

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}
	})

	t.Run("Contains on Queue with Poll", func(t *testing.T) {
		q := New[int]()
		q.Add(1)
		desiredValue := 1

		if !q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned false for queue: %s", desiredValue, q.String())
		}

		q.Poll()

		if q.Contains(desiredValue) {
			t.Errorf("\nq.Contains(%d) returned true for queue: %s", desiredValue, q.String())
		}
	})
}

func TestQueue_Clear(t *testing.T) {
	t.Run("Clear Should Return Queue with Size of 0", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		if q.Size() != 0 {
			t.Errorf("Size after clear is %d but expected 0", q.Size())
		}
	})

	t.Run("Clear then Poll Shuold Return nil", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		if q.Poll() != nil {
			t.Errorf("Poll after clear returned %d but expected nil", q.Peek())
		}
	})

	t.Run("Clear then Peek Should Return nil", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		if q.Peek() != nil {
			t.Errorf("Peek after clear returned %d but expected nil", q.Peek())
		}
	})

	t.Run("Clear then IsEmpty Should Return true", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7)

		q.Clear()

		if !q.IsEmpty() {
			t.Error("IsEmpty after clear returned false but expected true", q.IsEmpty())
		}
	})
}

func TestQueue_RemoveIf(t *testing.T) {
	t.Run("Queue Size Should Remain The Same When Filter Is Always False", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedSize := 9
		isGreaterThanTen := func(queueElement int) bool {
			return queueElement > 10
		}

		q.RemoveIf(isGreaterThanTen)

		if q.Size() != expectedSize {
			t.Errorf("Queue has incorrect size, expected %d but is %d", expectedSize, q.Size())
		}
	})

	t.Run("Queue Size Should Be 0 When Filter Removes All Elements", func(t *testing.T) {
		q := New[int]()
		q.Add(2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedSize := 0

		isLessThanOneHundred := func(queueElement int) bool {
			return queueElement < 100
		}

		q.RemoveIf(isLessThanOneHundred)

		if q.Size() != expectedSize {
			t.Errorf("Queue has incorrect size, expected %d but is %d", expectedSize, q.Size())
		}
	})

	t.Run("Queue Elements Should Not Change When Filter Is Always False", func(t *testing.T) {
		q := New[int]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedOrdering := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		isGreaterThanTen := func(queueElement int) bool {
			return queueElement > 10
		}

		q.RemoveIf(isGreaterThanTen)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Queue Should Not Contain Any Elements When Filter Is Always True", func(t *testing.T) {
		q := New[int]()
		q.Add(2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedOrdering := []int{}

		isLessThanOneHundred := func(queueElement int) bool {
			return queueElement < 100
		}

		q.RemoveIf(isLessThanOneHundred)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Error(msg)
		}
	})

	t.Run("Queue Size Should Decrease Properly", func(t *testing.T) {
		q := New[int]()
		q.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedSize := 12

		isGreaterThanTen := func(queueElement int) bool {
			return queueElement > 10
		}

		q.RemoveIf(isGreaterThanTen)

		actualSize := q.Size()

		if actualSize != expectedSize {
			t.Errorf("Queue has incorrect size, expected %d but is %d", expectedSize, actualSize)
		}
	})

	t.Run("Queue Nodes Should Be Properly Reorganized After RemoveIf Removes Several Elements", func(t *testing.T) {
		q := New[int]()
		q.Add(88, 2, 7, 3, 8, 1, 9, 1, 15, 3, 77, 66, 52, 5, 5, 5, 1)
		expectedOrdering := []int{2, 7, 3, 8, 1, 9, 1, 3, 5, 5, 5, 1}

		isGreaterThanTen := func(queueElement int) bool {
			return queueElement > 10
		}

		q.RemoveIf(isGreaterThanTen)

		valid, msg := validateQueue(expectedOrdering, *q)
		if !valid {
			t.Error(msg)
		}
	})
}

func TestQueue_Remove(t *testing.T) {
	t.Run("Initial test", func(t *testing.T) {
		q := New[int]()
		q.Add(2, 5, 3, 4, 1, 7)

		q.Remove(5)

		fmt.Println(q.String())
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("Size on Empty Queue", func(t *testing.T) {
		q := New[int64]()

		if q.Size() != 0 {
			t.Errorf("New queue has size of %d instead of 0!", q.Size())
		}
	})

	t.Run("Size on Non-Empty Queue", func(t *testing.T) {
		q := New[int64]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		if q.Size() != 10 {
			t.Errorf("Queue of 10 elements has size of %d!", q.Size())
		}
	})

	t.Run("Size of Queue After Adding and Polling", func(t *testing.T) {
		q := New[int64]()

		q.Add(1, 2, 3)
		if q.Size() != 3 {
			t.Errorf("Queue of size 3 has size of %d!", q.Size())
		}

		q.Poll()

		q.Add(4, 5, 6, 7, 8, 9)
		if q.Size() != 8 {
			t.Errorf("Queue of size 6 has size of %d!", q.Size())
		}

		for i := 0; i < 6; i++ {
			q.Poll()
		}

		if q.Size() != 2 {
			t.Errorf("Queue of size 2 has size of %d!", q.Size())
		}
	})
}

func TestQueue_IsEmpty(t *testing.T) {
	t.Run("IsEmpty on Empty Queue", func(t *testing.T) {
		q := New[int64]()

		if !q.IsEmpty() {
			t.Error("New queue is not empty!")
		}
	})

	t.Run("IsEmpty on Queue with Single Value", func(t *testing.T) {
		q := New[int64]()
		q.Add(1)

		if q.IsEmpty() {
			t.Error("Queue with single value is returning true for IsEmpty()!")
		}
	})

	t.Run("IsEmpty on Non-Empty Queue, then Empty Queue", func(t *testing.T) {
		q := New[int64]()
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
	t.Run("String on Empty Queue", func(t *testing.T) {
		q := New[int64]()
		if q.String() != "" {
			t.Error("Empty queue not returning empty string!")
		}
	})

	t.Run("String on Queue with Single Value", func(t *testing.T) {
		q := New[int64]()
		q.Add(1)

		expectedString := "1"
		actualString := q.String()

		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})

	t.Run("String on Non-Empty Queue", func(t *testing.T) {
		q := New[int64]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedString := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9"
		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})

	t.Run("String on Non-Empty Queue after Polling", func(t *testing.T) {
		q := New[int64]()
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		for i := 0; i < 5; i++ {
			q.Poll()
		}
		expectedString := "6 -> 7 -> 8 -> 9"
		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("\nExpected: %s\nGot: %s", expectedString, actualString)
		}
	})
}
