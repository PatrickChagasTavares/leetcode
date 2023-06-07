package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	cases := map[string]struct {
		expectedData []int
		inputNums    []int
		inputTarget  int
	}{
		"try target 9=[0,1]": {
			expectedData: []int{0, 1},
			inputNums:    []int{2, 7, 11, 15},
			inputTarget:  9,
		},
		"try target 6=[1,2]": {
			expectedData: []int{1, 2},
			inputNums:    []int{3, 2, 4},
			inputTarget:  6,
		},
		"try target 6=[0,1]": {
			expectedData: []int{0, 1},
			inputNums:    []int{3, 3},
			inputTarget:  6,
		},
		"try target 10=[4,5]": {
			expectedData: []int{4, 5},
			inputNums:    []int{3, 10, 22, 33, 44, -34, 1, 9},
			inputTarget:  10,
		},
		"try target 10=[0,2]": {
			expectedData: []int{0, 2},
			inputNums:    []int{3, 10, 7, 33, 44, -34, 1, 9},
			inputTarget:  10,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := twoSum(cs.inputNums, cs.inputTarget)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}
