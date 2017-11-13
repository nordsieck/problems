package atoms

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestAtom(t *testing.T) {
	defect.DeepEqual(t, Atom("H"), Result{map[string]int{`H`: 1}, ""})
	defect.DeepEqual(t, Atom("He2O"), Result{map[string]int{`He`: 2}, "O"})
}

func TestAdd(t *testing.T) {
	defect.DeepEqual(t, Add(), map[string]int{})
	defect.DeepEqual(t, Add(map[string]int{"H": 1}), map[string]int{"H": 1})
	defect.DeepEqual(t, Add(map[string]int{"H": 1}, map[string]int{"He": 2}), map[string]int{"H": 1, "He": 2})
	defect.DeepEqual(t, Add(map[string]int{"H": 1}, map[string]int{"H": 2}), map[string]int{"H": 3})
}

func TestParen(t *testing.T) {
	defect.DeepEqual(t, Paren(""), Result{})
	defect.DeepEqual(t, Paren("H"), Result{})
	defect.DeepEqual(t, Paren("(H)"), Result{freq: map[string]int{"H": 1}})
	defect.DeepEqual(t, Paren("(He2)"), Result{freq: map[string]int{"He": 2}})
	defect.DeepEqual(t, Paren("()"), Result{})
	defect.DeepEqual(t, Paren("()2"), Result{})
	defect.DeepEqual(t, Paren("(H)2"), Result{freq: map[string]int{"H": 2}})
	defect.DeepEqual(t, Paren("((H))"), Result{freq: map[string]int{"H": 1}})
	defect.DeepEqual(t, Paren("((H))2"), Result{freq: map[string]int{"H": 2}})
}

func TestProcess(t *testing.T) {
	defect.DeepEqual(t, Process("H"), map[string]int{"H": 1})
	defect.DeepEqual(t, Process("H2O"), map[string]int{"H": 2, "O": 1})
}

func TestCountOfAtoms(t *testing.T) {
	defect.Equal(t, countOfAtoms("H2O"), "H2O")
	defect.Equal(t, countOfAtoms("OH2"), "H2O")
	defect.Equal(t, countOfAtoms("Mg(OH)2"), "H2MgO2")
	defect.Equal(t, countOfAtoms("K4(ON(SO3)2)2"), "K4N2O14S4")
}
