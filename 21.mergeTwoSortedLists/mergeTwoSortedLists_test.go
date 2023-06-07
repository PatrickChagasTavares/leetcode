package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	cases := map[string]struct {
		expectedData *ListNode
		inputOne     *ListNode
		inputTwo     *ListNode
	}{
		"should return [1,1,2,3,4,4]": {
			expectedData: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
			inputOne: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			inputTwo: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		// "[1,2,4]
		// [1,3,4]
		// []
		// []
		// []
		// [0]"
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			resp := mergeTwoLists(cs.inputOne, cs.inputTwo)

			assert.Equal(t, cs.expectedData, resp)
		})
	}

}
