package longest

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestTrie_Insert(t *testing.T) {
	try := NewTrie()
	try.insert("a", 0)
	defect.DeepEqual(t, try, &trie{child: tm{
		"a"[0]: &trie{current: "a", child: tm{}},
	}, current: "head"})

	try.insert("ab", 0)
	defect.DeepEqual(t, try, &trie{child: tm{
		"a"[0]: &trie{current: "a", child: tm{
			"b"[0]: &trie{current: "ab", child: tm{}},
		}},
	}, current: "head"})
	try.insert("cc", 0)
	defect.DeepEqual(t, try, &trie{child: tm{
		"a"[0]: &trie{current: "a", child: tm{"b"[0]: &trie{current: "ab", child: tm{}}}},
		"c"[0]: &trie{child: tm{"c"[0]: &trie{current: "cc", child: tm{}}}},
	}, current: "head"})
}

func TestTrie_Longest(t *testing.T) {
	try := NewTrie()
	try.insert("a", 0)
	try.insert("ab", 0)
	try.insert("cc", 0)

	len, word := try.longest()
	defect.Equal(t, len, 2)
	defect.Equal(t, word, "ab")

	world := NewTrie()
	for _, s := range []string{"w", "wo", "wor", "worl", "world"} {
		world.insert(s, 0)
	}
	len, word = world.longest()
	defect.Equal(t, len, 5)
	defect.Equal(t, word, "world")

	fruit := NewTrie()
	for _, s := range []string{"a", "banana", "app", "appl", "ap", "apply", "apple"} {
		fruit.insert(s, 0)
	}
	len, word = fruit.longest()
	defect.Equal(t, len, 5)
	defect.Equal(t, word, "apple")
}

func TestLongestWord(t *testing.T) {
	defect.Equal(t, longestWord([]string{"w", "wo", "wor", "worl", "world"}), "world")
	defect.Equal(t, longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}), "apple")
}
