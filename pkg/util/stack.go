package util

import (
	"container/list"

	"github.com/mothdb-bd/common-go/pkg/basic"
)

type Stack[T basic.Object] struct {
	list *list.List
}

func NewStack[T basic.Object]() *Stack[T] {
	list := list.New()
	return &Stack[T]{
		list: list,
	}
}

func (p *Stack[T]) Push(v T) {
	p.list.PushBack(v)
}

func (p *Stack[T]) Pop() (T, bool) {
	e := p.list.Back()
	if e != nil {
		p.list.Remove(e)
		return e.Value.(T), true
	}
	return basic.NullT[T](), false
}

func (p *Stack[T]) Top() (T, bool) {
	e := p.list.Back()
	if e != nil {
		return e.Value.(T), true
	}
	return basic.NullT[T](), false
}

func (p *Stack[T]) Len() int {
	return p.list.Len()
}

func (p *Stack[T]) Empty() bool {
	return p.list.Len() == 0
}
