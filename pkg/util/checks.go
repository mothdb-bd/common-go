package util

import "github.com/mothdb-bd/common-go/pkg/basic"

// 校验
type Predicate[T basic.Object] interface {
	Test(t T) bool
}

type BiPredicate[T, U basic.Object] interface {
	Test(t T, u U) bool
}

type BiPredicateAdapter[T, U basic.Object] struct {
	predicate Predicate[T]
}

func NewBiPredicate[T, U basic.Object](predicate Predicate[T]) *BiPredicateAdapter[T, U] {
	adapter := &BiPredicateAdapter[T, U]{}
	adapter.predicate = predicate
	return adapter
}

func (adapter BiPredicateAdapter[T, U]) Test(t T, u U) bool {
	return adapter.predicate.Test(t)
}

func CheckState(expression bool) {
	if !expression {
		panic("Illegal state")
	}
}

func CheckState2(expression bool, msg string) {
	if !expression {
		panic(msg)
	}
}

func Verify(expression bool) {
	if !expression {
		panic("Verify")
	}
}

func Verify2(expression bool, msg string) {
	if !expression {
		panic(msg)
	}
}

func CheckArgument(expression bool) {
	if !expression {
		panic("Illegal state")
	}
}

func CheckArgument2(expression bool, msg string) {
	if !expression {
		panic(msg)
	}
}

func CheckNotNull[T basic.Object](reference T) T {
	if &reference == nil {
		panic("Null reference")
	}
	return reference
}
