package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	cases := map[string]struct {
		expectedData int
		inputNums    []int
		inputTarget  int
	}{
		"Should return position 2": {
			expectedData: 2,
			inputNums:    []int{1, 3, 5, 6},
			inputTarget:  5,
		},
		"Should return position 1": {
			expectedData: 1,
			inputNums:    []int{1, 3, 5, 6},
			inputTarget:  2,
		},
		"Should return position 4": {
			expectedData: 4,
			inputNums:    []int{1, 3, 5, 6},
			inputTarget:  7,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := searchInsertV2(cs.inputNums, cs.inputTarget)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}
