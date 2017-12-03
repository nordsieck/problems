package structures

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestStack_Contains(t *testing.T) {
	s := Stack{len: 2, head: &node{val: 5, next: &node{val: 6}}}
	defect.Equal(t, s.contains(3), false)
	defect.Equal(t, s.contains(6), true)
}

func TestStack_Delete(t *testing.T) {
	s := Stack{len: 2, head: &node{val: 5, next: &node{val: 6}}}
	s.delete()
	defect.DeepEqual(t, s, Stack{len: 1, head: &node{val: 6}})
}

func TestStack_Get(t *testing.T) {
	s := Stack{len: 2, head: &node{val: 5, next: &node{val: 6}}}
	defect.Equal(t, s.get(), 5)
}

func TestStack_Insert(t *testing.T) {
	s := Stack{len: 2, head: &node{val: 5, next: &node{val: 6}}}
	s.insert(4)
	defect.DeepEqual(t, s, Stack{len: 3, head: &node{val: 4, next: &node{val: 5, next: &node{val: 6}}}})
}
