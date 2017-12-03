package structures

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestHeap_Contains(t *testing.T) {
	defect.Equal(t, (&Heap{3, 4, 5}).contains(2), false)
	defect.Equal(t, (&Heap{3, 4, 5}).contains(4), true)
}

func TestHeap_Delete(t *testing.T) {
	h := Heap{5, 6, 7, 8, 9, 10}
	defect.Equal(t, h.delete(), 5)
	defect.DeepEqual(t, h, Heap{6, 8, 7, 10, 9})

	h = Heap{5}
	defect.Equal(t, h.delete(), 5)
	defect.DeepEqual(t, h, Heap{})
}

func TestHeap_Get(t *testing.T) {
	h := Heap{5, 6, 7, 8, 9, 10}
	defect.Equal(t, h.get(), 5)
	defect.DeepEqual(t, h, Heap{5, 6, 7, 8, 9, 10})
}

func TestHeap_Insert(t *testing.T) {
	h := Heap{5, 6, 7, 8, 9, 10}
	h.insert(11)
	defect.DeepEqual(t, h, Heap{5, 6, 7, 8, 9, 10, 11})

	h = Heap{5, 6, 7, 8, 9, 10}
	h.insert(4)
	defect.DeepEqual(t, h, Heap{4, 6, 5, 8, 9, 10, 7})
}

func TestHeap_Repair(t *testing.T) {
	h := Heap{5, 6, 7, 8, 9, 10}
	defect.Equal(t, h.repair(0), -1)
	defect.DeepEqual(t, h, Heap{5, 6, 7, 8, 9, 10})

	h = Heap{11, 6, 7, 8, 9, 10}
	defect.Equal(t, h.repair(0), 1)
	defect.DeepEqual(t, h, Heap{6, 11, 7, 8, 9, 10})

	defect.Equal(t, h.repair(1), 3)
	defect.DeepEqual(t, h, Heap{6, 8, 7, 11, 9, 10})
}

func TestHeap_Swap(t *testing.T) {
	h := Heap{5, 6, 7, 8, 9, 10}
	defect.Equal(t, h.swap(1, 3), false)
	defect.Equal(t, h.swap(2, 6), false)
	h = append(h, 1)
	defect.Equal(t, h.swap(2, 6), true)
	defect.DeepEqual(t, h, Heap{5, 6, 1, 8, 9, 10, 7})
}
