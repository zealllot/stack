package stack

import (
	"container/list"
	"errors"
)

type Stack struct {
	value list.List
}

func (s *Stack) Push(value interface{}) {
	s.value.PushBack(value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.value.Back() == nil {
		return nil, errors.New("The Stack is empty!")
	}
	return s.value.Remove(s.value.Back()), nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.value.Back() == nil {
		return nil, errors.New("The Stack is empty!")
	}
	return s.value.Back().Value, nil
}
