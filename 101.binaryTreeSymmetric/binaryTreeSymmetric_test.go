package binaryTreeSymmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "gotest.tools/assert"
)

func Test_isSymmetric(t *testing.T) {

	cases := map[string]struct {
		expectedData bool
		input        *TreeNode
	}{
		"expected: [1,2,2,3,4,4,3]=true": {
			expectedData: true,
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			data := isSymmetric(cs.input)

			assert.Equal(t, cs.expectedData, data)
		})
	}
}
