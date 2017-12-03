package structures

type Stack LinkedList

func (s *Stack) contains(val int) bool { return (*LinkedList)(s).contains(val) != nil }
func (s *Stack) delete()               { (*LinkedList)(s).delete(&s.head) }
func (s *Stack) get() int              { return (*(*LinkedList)(s).get(0)).val }
func (s *Stack) insert(val int)        { (*LinkedList)(s).insert(&s.head, val) }
