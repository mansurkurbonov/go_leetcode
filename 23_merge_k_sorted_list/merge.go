package mergeksortedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	numbers := []int{}

	for _, list := range lists {
		curr := list
		for curr != nil {
			numbers = append(numbers, curr.Val)
			curr = curr.Next
		}
	}
	if len(numbers) == 0 {
		return nil
	}
	numbers = bubbleSort(numbers)

	list := &ListNode{}
	for i, n := range numbers {
		if i == 0 {
			list.Val = n
		} else {
			newNode := ListNode{n, nil}
			iter := list
			for iter.Next != nil {
				iter = iter.Next
			}
			iter.Next = &newNode
		}
	}
	return list
}

func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
