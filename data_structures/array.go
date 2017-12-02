package structures

type Array []int

func (a *Array) contains(val int) int {
	for idx, i := range *a {
		if i == val {
			return idx
		}
	}
	return -1
}

func (a *Array) delete(idx int) {
	copy((*a)[idx:], (*a)[idx+1:])
	*a = (*a)[:len(*a)-1]
}

func (a *Array) get(idx int) int {
	if idx < 0 || idx > len(*a)-1 {
		return -1
	}
	return (*a)[idx]
}

func (a *Array) insert(idx, val int) {
	*a = append((*a)[:idx], append(Array{val}, (*a)[idx:]...)...)

}
