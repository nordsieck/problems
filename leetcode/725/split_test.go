package split

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestSizes(t *testing.T) {
	defect.DeepEqual(t, sizes(1, 1), []int{1})
	defect.DeepEqual(t, sizes(1, 2), []int{1, 0})
	defect.DeepEqual(t, sizes(1, 0), []int{})
	defect.DeepEqual(t, sizes(3, 5), []int{1, 1, 1, 0, 0})
	defect.DeepEqual(t, sizes(10, 3), []int{4, 3, 3})
}

func TestSplitListToParts(t *testing.T) {
	a := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	b := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}

	defect.DeepEqual(t, splitListToParts(nil, 5), []*ListNode{nil, nil, nil, nil, nil})
	defect.DeepEqual(t, splitListToParts(a, 5), []*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, nil, nil})

	defect.DeepEqual(t, splitListToParts(b, 3), []*ListNode{
		{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		{5, &ListNode{6, &ListNode{7, nil}}},
		{8, &ListNode{9, &ListNode{10, nil}}},
	})
}

func TestListNode_Len(t *testing.T) {
	a := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	b := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}

	defect.Equal(t, (*ListNode)(nil).len(), 0)
	defect.Equal(t, (&ListNode{}).len(), 1)
	defect.Equal(t, a.len(), 3)
	defect.Equal(t, b.len(), 10)
}
