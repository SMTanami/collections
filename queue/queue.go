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
	headAndTail := &node{}
	headAndTail.next = headAndTail
	return &queue{queueType: reflect.TypeOf(t), head: headAndTail, tail: headAndTail}
}

func (q *queue) Add(items ...any) error {
	for _, item := range items {
		isValid, err := q.validateType(item)
		if !isValid {
			return err
		}

		if q.size == 0 {
			q.head.val = item
		} else {
			q.tail.next = &node{val: item, next: nil}
			q.tail = q.tail.next
		}

		q.size++
	}

	return nil
}

func (q *queue) Peek() any {
	return q.head.val
}

func (q *queue) Poll() any {

	var returnValue any
	if q.size == 0 {
		return nil
	} else if q.size == 1 {
		returnValue = q.head.val
		q = New(q.queueType)
	} else {
		returnValue = q.head.val
		q.head = q.head.next
	}

	q.size--
	return returnValue
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func (q *queue) Print() {

	if q.size == 1 {
		fmt.Println(q.head.val)
	} else {
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

		fmt.Println(stringBuilder.String())
	}
}

func (q *queue) validateType(item any) (bool, error) {
	t := reflect.TypeOf(item)
	if t != q.queueType {
		return false, fmt.Errorf("%s invalid argument for queue of type %s", reflect.TypeOf(item).Name(), q.queueType.Name())
	}

	return true, nil
}
