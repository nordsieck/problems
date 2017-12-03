package structures

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestQueue_Contains(t *testing.T) {
	n := &node{val: 5, next: &node{val: 6}}
	q := Queue{tail: n.next, LinkedList: LinkedList{len: 2, head: n}}
	defect.DeepEqual(t, q.contains(4), false)
	defect.DeepEqual(t, q.contains(5), true)
}

func TestQueue_Delete(t *testing.T) {
	n := &node{val: 5, next: &node{val: 6}}
	q := Queue{tail: n.next, LinkedList: LinkedList{len: 2, head: n}}
	q.delete()

	o := &node{val: 6}
	defect.DeepEqual(t, q, Queue{tail: o, LinkedList: LinkedList{len: 1, head: o}})

	q.delete()
	defect.DeepEqual(t, q, Queue{})
}

func TestQueue_Get(t *testing.T) {
	n := &node{val: 5, next: &node{val: 6}}
	q := Queue{tail: n.next, LinkedList: LinkedList{len: 2, head: n}}
	defect.Equal(t, q.get(), 6)
	defect.Equal(t, (&Queue{}).get(), -1)
}

func TestQueue_Insert(t *testing.T) {
	n := &node{val: 5, next: &node{val: 6}}
	q := Queue{tail: n.next, LinkedList: LinkedList{len: 2, head: n}}
	q.insert(7)

	o := &node{val: 5, next: &node{val: 6, next: &node{val: 7}}}
	defect.DeepEqual(t, q, Queue{tail: o.next.next, LinkedList: LinkedList{head: o, len: 3}})

	q = Queue{}
	q.insert(3)

	o = &node{val: 3}
	defect.DeepEqual(t, q, Queue{tail: o, LinkedList: LinkedList{head: o, len: 1}})
}
