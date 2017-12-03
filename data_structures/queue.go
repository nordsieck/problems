package structures

type Queue struct {
	LinkedList
	tail *node
}

func (q *Queue) contains(val int) bool { return q.LinkedList.contains(val) != nil }

func (q *Queue) delete() {
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.len--
}

func (q *Queue) get() int {
	if q.tail == nil {
		return -1
	}
	return q.tail.val
}

func (q *Queue) insert(val int) {
	if q.tail == nil {
		q.tail = &node{val: val}
		q.head = q.tail
	} else {
		q.tail.next = &node{val: val}
		q.tail = q.tail.next
	}
	q.len++
}
