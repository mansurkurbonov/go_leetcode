package addtwonumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	s := []int{}
	curr := l1
	for curr != nil {
		s = append(s, curr.Val)
		curr = curr.Next
	}

	s2 := []int{}
	curr2 := l2
	for curr2 != nil {
		s2 = append(s2, curr2.Val)
		curr2 = curr2.Next
	}

	l := 0
	if len(s) >= len(s2) {
		l = len(s)
	} else {
		l = len(s2)
	}

	list := &ListNode{}

	lenA := len(s)
	lenB := len(s2)
	x := 0
	for i := 0; i < l; i++ {
		a := 0

		if lenA > 0 {
			a = s[i]
		} else {
			a = 0
		}
		lenA--

		b := 0
		if lenB > 0 {
			b = s2[i]
		} else {
			b = 0
		}

		lenB--

		c := a + b + x
		val := 0

		if c > 9 {
			x = (c / 10)
			val = c % 10

		} else {
			x = 0
			val = c
		}

		if i == 0 {
			list.Val = val
		} else {
			newNode := ListNode{val, nil}
			iter := list
			for iter.Next != nil {
				iter = iter.Next
			}
			iter.Next = &newNode
		}
		if i == l-1 && c > 9 {
			newNode := ListNode{x, nil}
			iter := list
			for iter.Next != nil {
				iter = iter.Next
			}
			iter.Next = &newNode
		}
	}

	return list
}
