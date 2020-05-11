package main

import "errors"

type Stack struct {
	List
}

func (s *Stack) Add(data interface{}) {
	n := &node{
		data: data,
	}

	if s.head == nil && s.tail == nil {
		s.head = n
		s.tail = n
	} else {
		n.prev = s.tail
		s.tail.next = n
		s.tail = n
	}
	s.len++
}

func (s *Stack) Get() (interface{}, error) {
	if s.len <= 0 {
		return nil, errors.New("empty stack")
	}
	n := s.tail
	if s.len == 1 {
		s.head = nil
		s.tail = nil
	} else {
		s.tail = s.tail.prev
	}
	s.len--
	return n.data, nil
}