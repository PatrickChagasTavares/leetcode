package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {

	cases := map[string]struct {
		expectedData int
		input        int
	}{
		"Should return 5": {
			expectedData: 5,
			input:        4,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := climbStairs(cs.input)

			assert.Equal(t, cs.expectedData, resp)
		})
	}
}
