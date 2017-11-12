package split

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) len() int {
	var ret int
	for root := l; root != nil; root = root.Next {
		ret++
	}
	return ret
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	ret := make([]*ListNode, 0, k)

	var len int
	for ptr := root; ptr != nil; ptr = ptr.Next {
		len++
	}

	for _, size := range sizes(len, k) {
		var next *ListNode
		next = root

		if size != 0 {
			for i := 0; i < size-1; i++ {
				root = root.Next
			}
			root, root.Next = root.Next, nil
		}
		ret = append(ret, next)
	}

	return ret
}

func sizes(len, num int) []int {
	ret := make([]int, 0, num)

	for num > 0 {
		if len/num*num == len {
			ret = append(ret, len/num)
			len -= len / num
		} else {
			ret = append(ret, len/num+1)
			len -= len/num + 1
		}
		num -= 1
	}
	return ret
}
