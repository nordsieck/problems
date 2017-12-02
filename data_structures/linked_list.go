package structures

type LinkedList struct {
	head *node
	len  int
}

type node struct {
	val  int
	next *node
}

func (l *LinkedList) contains(val int) **node {
	curr := &l.head
	for ; *curr != nil && (*curr).val != val; curr = &(*curr).next {
	}
	if *curr == nil {
		return nil
	}
	return curr
}

func (l *LinkedList) delete(n **node) {
	*n = (*n).next
	l.len--
}

func (l *LinkedList) get(idx int) **node {
	if idx >= l.len {
		return nil
	}
	curr := &l.head
	for i := 0; i < idx; i, curr = i+1, &(*curr).next {
	}
	return curr
}

func (l *LinkedList) insert(n **node, val int) {
	*n = &node{val: val, next: *n}
	l.len++
}
