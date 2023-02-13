package queue

import (
	"fmt"
	"reflect"
	"strings"
)

type node struct {
	val  any
	next *node
}

type queue struct {
	head      *node
	tail      *node
	size      int
	queueType reflect.Type
}

func New(t any) *queue {
	return &queue{queueType: reflect.TypeOf(t)}
}

func (q *queue) Add(items ...any) error {
	for _, item := range items {
		isValid, err := q.validateType(item)
		if !isValid {
			return err
		}

		if q.head == nil {
			initialNode := &node{val: item}
			q.head = initialNode
			q.tail = initialNode
		} else {
			q.tail.next = &node{val: item, next: nil}
			q.tail = q.tail.next
		}

		q.size++
	}

	return nil
}

func (q *queue) Peek() any {
	if q.head != nil {
		return q.head.val
	} else {
		return nil
	}
}

func (q *queue) Poll() any {
	var val any
	if q.size == 0 {
		return val
	} else {
		val = q.head.val
		q.head = q.head.next
	}
	q.size--
	return val
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func (q *queue) String() string {
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

func (q *queue) validateType(item any) (bool, error) {
	t := reflect.TypeOf(item)
	if t != q.queueType {
		return false, fmt.Errorf("%s invalid argument for queue of type %s", reflect.TypeOf(item).Name(), q.queueType.Name())
	}

	return true, nil
}
