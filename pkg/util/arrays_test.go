package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	src := []int{1, 1, 1, 1}
	dest := []int{0, 0, 0, 0}
	CopyArrays(src, 0, dest, 0, 4)

	for idx := range dest {
		assert.Equal(t, dest[idx], src[idx])
	}

	src[0] = 2
	assert.NotEqual(t, dest[0], src[0])
}

func TestCopy2(t *testing.T) {
	src := []int{1, 1, 1, 1}
	dest := []int{0, 0, 0, 0, 0, 0}
	assert.Equal(t, len(dest), 6)
	CopyArrays(src, 0, dest, 0, 4)

	for idx := range src {
		assert.Equal(t, dest[idx], src[idx])
	}
	assert.Equal(t, dest[4], 0)

	src1 := []int{1, 1, 1, 1}
	dest1 := []int{0, 0, 0, 0, 0, 0, 0}
	CopyArrays(src1, 0, dest1, 2, 4)
	assert.Equal(t, 0, dest1[0])
	assert.Equal(t, 0, dest1[1])
	assert.Equal(t, src1[0], dest1[2])
	assert.Equal(t, src1[1], dest1[3])
	assert.Equal(t, src1[2], dest1[4])
	assert.Equal(t, src1[3], dest1[5])
	assert.Equal(t, 0, dest1[6])

}
