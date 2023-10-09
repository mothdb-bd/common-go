package optional

import (
	"strconv"
)

type OptionalLong struct {
	value    int64
	hasValue bool
}

func OptionalLongEmpty() *OptionalLong {
	o := new(OptionalLong)
	o.hasValue = false
	return o
}

func OptionalLongOf(value int64) *OptionalLong {
	ot := new(OptionalLong)
	ot.value = value
	ot.hasValue = true
	return ot
}

func (op *OptionalLong) String() string {
	if op.hasValue {
		return strconv.FormatInt(op.value, 10)
	} else {
		return "empty"
	}
}

func (op *OptionalLong) Get() int64 {
	return op.value
}

func (op *OptionalLong) IsEmpty() bool {
	return !op.hasValue
}

func (op *OptionalLong) IsPresent() bool {
	return op.hasValue
}

func (op *OptionalLong) Equals(value int64) bool {
	// Special case: value type is slice(map, struct)
	if op.hasValue {
		return value == op.value
	}
	return false
}

func (op *OptionalLong) OrElse(other int64) int64 {
	if op.hasValue {
		return op.value
	} else {
		return other
	}
}
