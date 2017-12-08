package numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	one, two := l1, l2
	var head *ListNode
	var curr **ListNode = &head

	for ; one != nil && two != nil; one, two, curr = one.Next, two.Next, &(*curr).Next {
		*curr = &ListNode{Val: (one.Val + two.Val + carry) % 10}
		carry = (one.Val + two.Val + carry) / 10
	}

	if one == nil {
		one, two = two, one
	}
	for ; one != nil; one, curr = one.Next, &(*curr).Next {
		*curr = &ListNode{Val: (one.Val + carry) % 10}
		carry = (one.Val + carry) / 10
	}

	if carry == 1 {
		*curr = &ListNode{Val: carry}
	}

	return head
}
