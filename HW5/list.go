package main

import "errors"

type List struct {
	head *node
	tail *node
	len  int
}

type node struct {
	data interface{}
	next *node
	prev *node
}

func (l *List) Add(data interface{}) {
	n := &node{
		data: data,
	}

	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.len++
}

func (l *List) get(i int) (*node, error) {
	if i < 0 || i >= l.len {
		return nil, errors.New("out of range")
	}
	index := 0

	current := l.head
	for index != i {
		current = current.next
		index++
	}
	return current, nil
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Get(i int) (interface{}, error) {
	node, err := l.get(i)
	if err != nil {
		return nil, err
	}
	return node.data, nil
}

func (l *List) Remove(i int) error {
	node, err := l.get(i)
	if err != nil {
		return err
	}

	if node == l.head {
		l.head = node.next
	}

	if node == l.tail {
		l.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	return nil
}
