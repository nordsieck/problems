package decode

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestDecodeString(t *testing.T) {
	defect.Equal(t, decodeString(`3[a]4[b]`), `aaabbbb`)
	defect.Equal(t, decodeString(`3[a]2[bc]`), `aaabcbc`)
	defect.Equal(t, decodeString(`3[a2[c]]`), `accaccacc`)
	defect.Equal(t, decodeString(`2[abc]3[cd]ef`), `abcabccdcdcdef`)
}

func TestParse(t *testing.T) {
	defect.Equal(t, parse(`a`), `a`)
	defect.Equal(t, parse(`abc`), `abc`)
	defect.Equal(t, parse(`2[a]`), `aa`)
	defect.Equal(t, parse(`2[3[a]]`), `aaaaaa`)
	defect.Equal(t, parse(`a2[b]c`), `abbc`)
}
