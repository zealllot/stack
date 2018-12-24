package stack

import (
	"container/list"
	"errors"
	"fmt"
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

func (s *Stack) Length() int {
	return s.value.Len()
}

func (s *Stack)Print()  {
	var here=s.value.Back()
	fmt.Println(here.Value)
	for (here.Prev()!=nil){
		here=here.Prev()
		fmt.Println(here.Value)
	}
}
