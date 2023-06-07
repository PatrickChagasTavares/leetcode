package stockbuysell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxProfit(t *testing.T) {
	cases := map[string]struct {
		expectedData int
		input        []int
	}{
		"expected: 5 to [7,1,5,3,6,4]": {
			expectedData: 5,
			input:        []int{7, 1, 5, 3, 6, 4},
		},
		"expected: 0 to [7,6,4,3,1]": {
			expectedData: 0,
			input:        []int{7, 6, 4, 3, 1},
		},
		"expected: 2 to [2,4,1]": {
			expectedData: 2,
			input:        []int{2, 4, 1},
		},
		"expected: 2 to [2,1,2,1,0,1,2]": {
			expectedData: 2,
			input:        []int{2, 1, 2, 1, 0, 1, 2},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			data := maxProfit(cs.input)

			assert.Equal(t, cs.expectedData, data)
		})
	}
}
