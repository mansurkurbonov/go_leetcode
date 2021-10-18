package addtwonumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	in := []struct {
		l1  *ListNode
		l2  *ListNode
		out *ListNode
	}{
		{
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			out: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val:  0,
				Next: nil,
			},
			l2: &ListNode{
				Val:  0,
				Next: nil,
			},
			out: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	t.Run("1 test", func(t *testing.T) {
		output := addTwoNumbers2(in[0].l1, in[0].l2)
		if output.Val != in[0].out.Val && output.Next.Val != in[0].out.Next.Val && output.Next.Next.Val != in[0].out.Next.Next.Val {
			t.Error("got", output, "expect", in[0].out)
		}
	})

	t.Run("2 test", func(t *testing.T) {
		output := addTwoNumbers2(in[1].l1, in[1].l2)
		if output.Val != in[1].out.Val {
			t.Error("got", output, "expect", in[0].out)
		}
	})

}
