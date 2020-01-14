package container

import (
	l "container/list"
	"sync"
)

type Stack struct {
	list *l.List
	mu sync.Mutex
}

func NewStack() *Stack {
	list := l.New()
	return &Stack{list: list,}
}

func (s *Stack) Push(t interface{}){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.list.PushFront(t)
}

func  (s *Stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	ele := s.list.Front()
	if nil != ele {
		s.list.Remove(ele)
		return ele.Value
	}

	return nil
}

func (s *Stack) Peak() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	ele := s.list.Front()
	return ele.Value
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
