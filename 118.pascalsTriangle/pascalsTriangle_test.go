package pascalstriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	cases := map[string]struct {
		expectedData [][]int
		input        int
	}{
		"expected: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]] to 5": {
			input: 5,
			expectedData: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		"expected: [[1]] to 1": {
			input:        1,
			expectedData: [][]int{{1}},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			data := generate(cs.input)

			assert.Equal(t, cs.expectedData, data)
		})
	}

}
