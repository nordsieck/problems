package structures

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestTrie_Contains(t *testing.T) {
	trie := &Trie{children: map[byte]*Trie{
		'a': &Trie{leaf: true, children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				's': &Trie{leaf: true},
				'p': &Trie{leaf: true},
			}},
		}},
	}}

	defect.Equal(t, trie.contains("asp"), true)
	defect.Equal(t, trie.contains("foo"), false)
	defect.Equal(t, trie.contains("as"), false)
	defect.Equal(t, trie.contains("a"), true)
}

func TestTrie_Delete(t *testing.T) {
	trie := &Trie{children: map[byte]*Trie{
		'a': &Trie{leaf: true, children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				's': &Trie{leaf: true},
				'p': &Trie{leaf: true},
			}},
		}},
	}}

	defect.Equal(t, trie.children['a'].delete(``), false)
	defect.DeepEqual(t, trie, &Trie{children: map[byte]*Trie{
		'a': &Trie{children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				's': &Trie{leaf: true},
				'p': &Trie{leaf: true},
			}},
		}},
	}})

	defect.Equal(t, trie.delete(`ass`), false)
	defect.DeepEqual(t, trie, &Trie{children: map[byte]*Trie{
		'a': &Trie{children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				'p': &Trie{leaf: true},
			}},
		}},
	}})

	defect.Equal(t, trie.delete(`foo`), false)
	defect.DeepEqual(t, trie, &Trie{children: map[byte]*Trie{
		'a': &Trie{children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				'p': &Trie{leaf: true},
			}},
		}},
	}})

	defect.Equal(t, trie.delete(`asp`), true)
	defect.DeepEqual(t, trie, &Trie{children: map[byte]*Trie{}})
}

func TestTrie_get(t *testing.T) {
	trie := &Trie{children: map[byte]*Trie{
		'a': &Trie{leaf: true, children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				's': &Trie{leaf: true, children: map[byte]*Trie{}},
				'p': &Trie{leaf: true, children: map[byte]*Trie{}},
			}},
		}},
	}}

	defect.Equal(t, trie.get(`ass`), trie.children['a'].children['s'].children['s'])
}

func TestTrie_Insert(t *testing.T) {
	trie := &Trie{children: map[byte]*Trie{
		'a': &Trie{leaf: true, children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				's': &Trie{leaf: true, children: map[byte]*Trie{}},
				'p': &Trie{leaf: true, children: map[byte]*Trie{}},
			}},
		}},
	}}

	trie.insert(`at`)
	defect.DeepEqual(t, trie, &Trie{children: map[byte]*Trie{
		'a': &Trie{leaf: true, children: map[byte]*Trie{
			's': &Trie{children: map[byte]*Trie{
				's': &Trie{leaf: true, children: map[byte]*Trie{}},
				'p': &Trie{leaf: true, children: map[byte]*Trie{}},
			}},
			't': &Trie{leaf: true, children: map[byte]*Trie{}},
		}},
	}})

	trie = &Trie{children: map[byte]*Trie{}}
	trie.insert(`at`)
	defect.DeepEqual(t, trie, &Trie{children: map[byte]*Trie{
		'a': &Trie{children: map[byte]*Trie{
			't': &Trie{leaf: true, children: map[byte]*Trie{}},
		}},
	}})
}
