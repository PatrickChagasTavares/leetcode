package binarytreemaxdepth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxDepth(t *testing.T) {

	cases := map[string]struct {
		expectedData int
		input        *TreeNode
	}{
		"expected: 3 to [3,9,20,null,null,15,7]": {
			expectedData: 3,
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 15,
					},
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
		"expected: 2 to [1,null,2]": {
			expectedData: 2,
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			data := maxDepth(cs.input)

			assert.Equal(t, cs.expectedData, data)
		})
	}
}
