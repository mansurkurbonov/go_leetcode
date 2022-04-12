package mergeksortedlist

import (
	"strconv"
	"testing"
)

func TestMergeKList(t *testing.T) {
	in := []struct {
		lists  []*ListNode
		expect *ListNode
	}{
		{
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
				{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
			},
			expect: &ListNode{
				Val: 1,
				Next: &ListNode{
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
										Next: &ListNode{
											Val: 5,
											Next: &ListNode{
												Val:  6,
												Next: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for i := 0; i < len(in); i++ {
		t.Run("test № "+strconv.Itoa(i+1), func(t *testing.T) {
			output := mergeKLists(in[i].lists)
			if output.Val != in[i].expect.Val && output.Next.Val != in[i].expect.Next.Val {
				t.Error("got", output, "expect", in[i].expect)
			}
		})
	}
}
