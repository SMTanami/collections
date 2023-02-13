package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Add(t *testing.T) {
	t.Run("Add One to Empty Queue", func(t *testing.T) {
		q := New(1)
		q.Add(1)
		if q.IsEmpty() {
			t.Error("Failed to add element to queue!")
		}
	})

	t.Run("Add Many to Empty Queue", func(t *testing.T) {
		q := New(1)

		q.Add(1, 2, 3, 4)

		var val any
		for i := 1; i <= q.size; i++ {
			val = q.Poll()
			if val != i {
				t.Errorf("Failed to add element to queue! Expected %d but got %d", i, val)
			}
		}
	})

	t.Run("Add Improper Value to Queue", func(t *testing.T) {
		q := New("Hello!")
		if q.Add(1) == nil {
			t.Error("Failed to ensure proper queue type! Added element of type int to queue of type string")
		}
	})
}

func TestQueue_Poll(t *testing.T) {
	t.Run("Poll from Queue", func(t *testing.T) {
		q := New(1)
		q.Add(1)
		var val any = q.Poll()

		if val != 1 {
			t.Errorf("Poll failed! Expected 1 but got %v", val)
		}
	})

	t.Run("Poll Head, then Add Elements and Poll Them", func(t *testing.T) {
		q := New(1)

		q.Add(1)
		q.Poll()

		q.Add(1, 2, 3, 4)
		var val any
		for i := 1; i <= q.size; i++ {
			val = q.Poll()
			if val != i {
				t.Errorf("Polling from empty queue resulted in ordering error! Expected %d but got %d\nQueue: %s", i, val, q.String())
			}
		}
	})

	t.Run("Poll from Empty Queue", func(t *testing.T) {
		q := New(1)

		for i := 0; i < 5; i++ {
			q.Poll()
		}

		q.Add(1, 2, 3, 4)
		var val any
		for i := 1; i <= q.size; i++ {
			val = q.Poll()
			if val != i {
				t.Errorf("Polling from empty queue resulted in ordering error! Expected %d but got %d\nQueue: %s", i, val, q.String())
			}
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("Peek on Empty Queue", func(t *testing.T) {
		q := New(1)

		val := q.Peek()
		if val != nil {
			t.Errorf("Peek on empty queue resulted in a non-nil value, %v!", val)
		}
	})

	t.Run("Peek on Queue with Single Value", func(t *testing.T) {
		q := New(1)
		q.Add(1)

		val := q.Peek()
		if val != 1 {
			t.Errorf("Peek on queue resulted in an unexpected value, %v!", val)
		}
	})

	t.Run("Peek on Queue with Multiple Values", func(t *testing.T) {
		q := New(1)
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)

		for i := 1; i <= q.size; i++ {
			val := q.Peek()
			if val != i {
				t.Errorf("Peek on queue resulted in an unexpected value, %v!", val)
				fmt.Printf("i=%d, peek=%d\n", i, val)
			}
			q.Poll()
		}

	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("Size on Empty Queue", func(t *testing.T) {
		q := New(1)

		if q.Size() != 0 {
			t.Errorf("New queue has size of %d instead of 0!", q.Size())
		}
	})

	t.Run("Size on Non-Empty Queue", func(t *testing.T) {
		q := New(1)
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		if q.Size() != 10 {
			t.Errorf("Queue of 10 elements has size of %d!", q.Size())
		}
	})

	t.Run("Size of Queue After Adding and Polling", func(t *testing.T) {
		q := New(1)

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
		q := New(1)

		if !q.IsEmpty() {
			t.Error("New queue is not empty!")
		}
	})

	t.Run("IsEmpty on Queue with Single Value", func(t *testing.T) {
		q := New(1)
		q.Add(1)

		if q.IsEmpty() {
			t.Error("Queue with single value is returning true for IsEmpty()!")
		}
	})

	t.Run("IsEmpty on Non-Empty Queue, then Empty Queue", func(t *testing.T) {
		q := New(1)
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
		q := New(1)
		if q.String() != "" {
			t.Error("Empty queue not returning empty string!")
		}
	})

	t.Run("String on Queue with Single Value", func(t *testing.T) {
		q := New(1)
		q.Add(1)

		expectedString := "1"
		actualString := q.String()

		if actualString != expectedString {
			t.Errorf("\nReturned: %s\nGot: %s", actualString, expectedString)
		}
	})

	t.Run("String on Non-Empty Queue", func(t *testing.T) {
		q := New(1)
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		expectedString := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9"
		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("\nReturned: %s\nGot: %s", actualString, expectedString)
		}
	})

	t.Run("String on Non-Empty Queue after Polling", func(t *testing.T) {
		q := New(1)
		q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
		for i := 0; i < 5; i++ {
			q.Poll()
		}
		expectedString := "6 -> 7 -> 8 -> 9"
		actualString := q.String()
		if actualString != expectedString {
			t.Errorf("\nReturned: %s\nGot: %s", actualString, expectedString)
		}
	})
}
