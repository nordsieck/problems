package structures

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestArray_Contains(t *testing.T) {
	a := Array{3, 4, 5}
	defect.Equal(t, a.contains(2), -1)
	defect.Equal(t, a.contains(3), 0)
}

func TestArray_Delete(t *testing.T) {
	a := Array{3, 4, 5}
	a.delete(1)
	defect.DeepEqual(t, a, Array{3, 5})
}

func TestArray_Get(t *testing.T) {
	a := Array{3, 4, 5}
	defect.Equal(t, a.get(1), 4)
}

func TestArray_Insert(t *testing.T) {
	a := Array{3, 4, 5}
	a.insert(1, 6)
	defect.DeepEqual(t, a, Array{3, 6, 4, 5})
	a.insert(0, 7)
	defect.DeepEqual(t, a, Array{7, 3, 6, 4, 5})

}
