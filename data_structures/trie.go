package structures

type Trie struct {
	leaf     bool
	children map[byte]*Trie
}

func (t *Trie) contains(val string) bool {
	if val == "" {
		return t.leaf
	}
	if _, ok := t.children[val[0]]; !ok {
		return false
	}
	return t.children[val[0]].contains(val[1:])
}

func (t *Trie) delete(val string) bool {
	if val == "" {
		if len(t.children) != 0 {
			t.leaf = false
			return false
		} else {
			return true
		}
	}
	if _, ok := t.children[val[0]]; !ok {
		return false
	}

	if t.children[val[0]].delete(val[1:]) {
		delete(t.children, val[0])
		return len(t.children) == 0
	}

	return false
}

// doesn't really make sense, but might if there were other fields
// in Trie
func (t *Trie) get(s string) *Trie {
	if s == `` {
		if t.leaf {
			return t
		}
		return nil
	}
	if _, ok := t.children[s[0]]; !ok {
		return nil
	}
	return t.children[s[0]].get(s[1:])
}

func (t *Trie) insert(s string) {
	if s == `` {
		t.leaf = true
		return
	}
	if _, ok := t.children[s[0]]; !ok {
		t.children[s[0]] = &Trie{children: map[byte]*Trie{}}
	}
	t.children[s[0]].insert(s[1:])
}
