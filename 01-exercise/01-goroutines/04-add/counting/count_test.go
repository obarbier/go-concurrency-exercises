package counting

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func BenchmarkAdd(b *testing.B) {
	numbers := GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	numbers := GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		AddConcurrent(numbers)
	}
}


func TestSplitSlice(t *testing.T) {
	// EVEN NUMBER
	var s []int
	var chunk int
	// TEST-1: EVEN NUMBER
	s = []int{1,2,3,4,5,6,7,8}
	chunk=4
	s1 := splitSlice(s, chunk)
	assert.Equal(t, chunk ,len(s1))
	assert.Equal(t,[][]int{ {1,2},{3,4}, {5,6}, {7,8}}, s1)

	// TEST-2: EVEN NUMBER
	s = []int{1,2,3,4,5,6,7,8}
	chunk=3
	s2 := splitSlice(s, chunk)
	fmt.Println(s2)
	assert.Equal(t, chunk , len(s2))
	assert.Equal(t,[][]int{ {1,2,3},{4,5,6}, {7,8}} , s2)

	// TEST-3: ODD NUMBER
	s = []int{1,2,3,4,5,6,7,8 ,9}
	chunk=4
	s3 := splitSlice(s, chunk)
	assert.Equal(t, chunk ,len(s3))
	assert.Equal(t,[][]int{ {1,2},{3,4}, {5,6}, {7,8,9}}, s3)

	// TEST-4: ODD NUMBER
	s = []int{1,2,3,4,5,6,7,8 ,9}
	chunk=3
	s4 := splitSlice(s, chunk)
	fmt.Println(s4)
	assert.Equal(t, chunk , len(s4))
	assert.Equal(t,[][]int{ {1,2,3},{4,5,6}, {7,8,9}} , s4)
}
