package optional

import (
	"fmt"
)

// OptionalEmpty() Optional --- return an empty Optional instance
// OptionalOf(value) Optional --- return an Optional with the specified present non-null value
// Optional.ToString() string --- return a non-empty string representation of this optional suitable for debugging
// Optional.Get() interface{} --- If a value is present in this Optional, returns the value
// Optional.IsPresent() bool --- Whether there is a value in Optional
// Optional.Equal(obj interface{}) bool --- Indicates whether some other object is "equal to" this Optional
type OptionalDouble struct {
	data     float64
	hasValue bool
}

// How does go implement static method
// - no static method, use function instead
func OptionalDoubleEmpty() *OptionalDouble {
	ot := new(OptionalDouble)
	ot.hasValue = false
	return ot
}

func OptionalDoubleOf(value float64) *OptionalDouble {
	ot := new(OptionalDouble)
	ot.data = value
	ot.hasValue = true
	return ot
}

func (op *OptionalDouble) String() string {
	if op.hasValue {
		str := fmt.Sprintf("%v", op.data)
		return str
	} else {
		return "empty"
	}
}

func (op *OptionalDouble) Get() float64 {
	return op.data
}

func (op *OptionalDouble) IsPresent() bool {
	return op.hasValue
}

func (op *OptionalDouble) IsEmpty() bool {
	return !op.hasValue
}

func (op *OptionalDouble) Equals(value float64) bool {
	// Special case: value type is slice(map, struct)
	if op.hasValue {
		return value == op.data
	}
	return false
}

func (op *OptionalDouble) OrElse(other float64) float64 {
	if op.hasValue {
		return op.data
	} else {
		return other
	}
}

// Could I use "[]" for custom type
// - maybe answer can be found in source code of map
