package structures

type Heap []int

func (h *Heap) contains(val int) bool {
	for _, i := range *h {
		if i == val {
			return true
		}
	}
	return false
}

func (h *Heap) delete() int {
	ret := (*h)[0]
	(*h)[0], *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]

	for idx := h.repair(0); idx != -1; idx = h.repair(idx) {
	}

	return ret
}

func (h *Heap) repair(idx int) int {
	var other int
	switch {
	case idx*2+2 < len(*h):
		if (*h)[idx*2+2] < (*h)[idx*2+1] {
			other = idx*2 + 2
		} else {
			other = idx*2 + 1
		}
	case idx*2+2 == len(*h):
		other = idx*2 + 1
	case idx*2+1 >= len(*h):
		return -1
	}
	if h.swap(idx, other) {
		return other
	}
	return -1
}

func (h *Heap) get() int { return (*h)[0] }

func (h *Heap) insert(i int) {
	*h = append(*h, i)

	for idx, next := len(*h)-1, len(*h)/2-1; (*h)[idx] < (*h)[next]; idx, next = next, (next-1)/2 {
		(*h)[next], (*h)[idx] = (*h)[idx], (*h)[next]
	}
}

func (h *Heap) swap(i, j int) bool {
	if j < len(*h) && (*h)[j] < (*h)[i] {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		return true
	}
	return false
}
