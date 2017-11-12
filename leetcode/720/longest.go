package longest

func longestWord(words []string) string {
	try := NewTrie()
	for _, w := range words {
		try.insert(w, 0)
	}
	_, word := try.longest()
	return word
}

type tm map[byte]*trie

type trie struct {
	child   tm
	current string
}

func NewTrie() *trie { return &trie{child: tm{}, current: "head"} }

func (t *trie) insert(word string, depth int) {
	if depth == len(word) {
		t.current = word
		return
	}

	if _, ok := t.child[word[depth]]; !ok {
		t.child[word[depth]] = &trie{child: tm{}}
	}
	t.child[word[depth]].insert(word, depth+1)
}

func (t *trie) longest() (length int, word string) {
	if t.current == "" {
		return -2, ""
	}
	length = 0
	word = t.current
	for _, try := range t.child {
		l, w := try.longest()
		if l+1 > length || (l+1 == length && w < word) {
			length = l + 1
			word = w
		}
	}
	return length, word
}
