package binarytreeinorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {

	cases := map[string]struct {
		expectedData []int
		input        *TreeNode
	}{
		"expected: [1,3,2]": {
			expectedData: []int{3, 4, 1, 1, 5},
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			data := inorderTraversal(cs.input)

			assert.Equal(t, cs.expectedData, data)
		})
	}
}
