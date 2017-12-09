package numbers

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestAddTwoNumbers(t *testing.T) {
	defect.DeepEqual(t, addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}),
		&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}})
	defect.DeepEqual(t, addTwoNumbers(&ListNode{Val: 5}, &ListNode{Val: 5}), &ListNode{Val: 0, Next: &ListNode{Val: 1}})
}
