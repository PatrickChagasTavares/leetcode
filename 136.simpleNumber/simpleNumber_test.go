package simplenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simpleNumber(t *testing.T) {

	cases := map[string]struct {
		inputNums    []int
		expectedData int
	}{
		"Should return 1 to input: [2,2,1]": {
			expectedData: 1,
			inputNums:    []int{2, 2, 1},
		},
		"Should return 4 to input: [4,1,2,1,2]": {
			expectedData: 4,
			inputNums:    []int{4, 1, 2, 1, 2},
		},
		"Should return 1 to input: [1]": {
			expectedData: 1,
			inputNums:    []int{1},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			data := singleNumber(cs.inputNums)

			assert.Equal(t, cs.expectedData, data)
		})
	}
	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			data := singleNumberV2(cs.inputNums)

			assert.Equal(t, cs.expectedData, data)
		})
	}
}

func Benchmark_singleNumber(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data := singleNumber(nums)
		if data != 4 {
			b.Error(data)
		}
	}
}
func Benchmark_singleNumberV2(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data := singleNumberV2(nums)
		if data != 4 {
			b.Error(data)
		}
	}
}
