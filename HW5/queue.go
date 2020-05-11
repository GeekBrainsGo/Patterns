package main

import "errors"

type Queue struct {
	List
}

func (q *Queue) Add(data interface{}) {
	n := &node{
		data: data,
	}

	if q.head == nil && q.tail == nil {
		q.head = n
		q.tail = n
	} else {
		n.prev = q.tail
		q.tail.next = n
		q.tail = n
	}
	q.len++
}

func (q *Queue) Get() (interface{}, error) {
	if q.len <= 0 {
		return nil, errors.New("empty queue")
	}
	n := q.head
	if q.len == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.len--
	return n.data, nil
}
