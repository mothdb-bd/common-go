package util

import (
	"testing"
)

func TestSetNonTSNew(t *testing.T) {
	set := NewSet[int](SET_NonThreadSafe)

	set.Add(2, 1)

	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := set.Has(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Has(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Has(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetAdd(t *testing.T) {
	set := NewSet[int](SET_NonThreadSafe)
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestSetHas(t *testing.T) {
	set := NewSet[int](SET_NonThreadSafe)
	set.Add(3, 1, 2)
	set.Add(2, 3)
	set.Add()
	// TODO 如果空时？
	// if actualValue := set.Has(); actualValue != true {
	// 	t.Errorf("Got %v expected %v", actualValue, true)
	// }
	if actualValue := set.Has(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Has(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Has(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[int](SET_NonThreadSafe)
	set.Add(3, 1, 2)
	set.Remove()
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	set.Remove(1)
	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	set.Remove(3)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if actualValue := set.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

// func TestSetString(t *testing.T) {
// 	c := newNonTS[int]()
// 	c.Add(1)
// 	if !strings.HasPrefix(c.String(), "HashSet") {
// 		t.Errorf("String should start with container name")
// 	}
// }

func TestSetIsSubset(t *testing.T) {
	set := NewSet[string](SET_NonThreadSafe)
	another := NewSet[string](SET_NonThreadSafe)

	if actualValue := set.IsSubset(another); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	if actualValue := set.IsSubset(another); actualValue == true {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

// TODO fixme SET用此方法无法编译通过
func TestSetUnion(t *testing.T) {
	set := NewSet[string](SET_NonThreadSafe)
	another := NewSet[string](SET_NonThreadSafe)

	union := Union(set, another)
	if actualValue, expectedValue := union.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	union = Union(set, another)

	if actualValue, expectedValue := union.Size(), 6; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := union.Has("a", "b", "c", "d", "e", "f"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetDifference(t *testing.T) {
	set := NewSet[string](SET_NonThreadSafe)
	another := NewSet[string](SET_NonThreadSafe)

	difference := Difference(set, another)
	if actualValue, expectedValue := difference.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	difference = Difference(set, another)

	if actualValue, expectedValue := difference.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := difference.Has("a", "b"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestIntersection(t *testing.T) {
	set := NewSet[string](SET_NonThreadSafe)
	another := NewSet[string](SET_NonThreadSafe)

	intersection := Intersection(set, another)
	if actualValue, expectedValue := intersection.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	intersection = Intersection(set, another)

	if actualValue, expectedValue := intersection.Size(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := intersection.Has("c", "d"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSymmetricDifference(t *testing.T) {
	set := NewSet[string](SET_NonThreadSafe)
	another := NewSet[string](SET_NonThreadSafe)

	symmetricdifference := SymmetricDifference(set, another)
	if actualValue, expectedValue := symmetricdifference.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	symmetricdifference = SymmetricDifference(set, another)

	if actualValue, expectedValue := symmetricdifference.Size(), 4; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := symmetricdifference.Has("a", "b", "e", "f"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func benchmarkHas(b *testing.B, set SetInterface[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Has(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set SetInterface[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set SetInterface[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkTSSetHas100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkNOTSSetHas100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkTSSetHas1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkNOTSSetHas1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkTSSetHas10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkNOTSSetHas10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkTSSetHas100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkNOTSSetHas100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkHas(b, set, size)
}

func BenchmarkTSSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewSet[int](SET_ThreadSafe)
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkNOTSSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewSet[int](SET_NonThreadSafe)
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTSSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkNOTSSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTSSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkNOTSSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTSSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkNOTSSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTSSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkNOTSSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTSSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkNOTSSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTSSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkNOTSSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTSSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewSet[int](SET_ThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkNOTSSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewSet[int](SET_NonThreadSafe)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}
