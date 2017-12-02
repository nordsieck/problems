package structures

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestLinkedList_Add(t *testing.T) {
	l := LinkedList{len: 2, head: &node{val: 2, next: &node{val: 3}}}
	defect.Equal(t, l.get(1), &l.head.next)
}

func TestLinkedList_Delete(t *testing.T) {
	l := LinkedList{len: 2, head: &node{val: 2, next: &node{val: 3}}}
	defect.Equal(t, l.get(1), &l.head.next)
	defect.Equal(t, l.len, 2)

	l.delete(&l.head.next)
	defect.Equal(t, l.get(1), (**node)(nil))
	defect.Equal(t, l.len, 1)
}

func TestLinkedList_Contains(t *testing.T) {
	l := LinkedList{len: 2, head: &node{val: 2, next: &node{val: 3}}}
	defect.Equal(t, l.contains(1), (**node)(nil))
	defect.Equal(t, l.contains(2), &l.head)
	defect.Equal(t, l.contains(3), &l.head.next)
	defect.Equal(t, l.contains(4), (**node)(nil))
}

func TestLinkedList_Get(t *testing.T) {
	l := LinkedList{}
	defect.Equal(t, l.get(1), (**node)(nil))

	l = LinkedList{len: 2, head: &node{val: 2, next: &node{val: 3}}}
	defect.Equal(t, l.get(1), &l.head.next)
}

func TestLinkedList_Insert(t *testing.T) {
	l := LinkedList{len: 2, head: &node{val: 2, next: &node{val: 3}}}
	l.insert(&l.head.next, 4)

	defect.DeepEqual(t, l, LinkedList{len: 3, head: &node{val: 2, next: &node{val: 4, next: &node{val: 3}}}})
}
