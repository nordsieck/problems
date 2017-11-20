package longest

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestLongestSubstring(t *testing.T) {
	defect.Equal(t, longestSubstring("aaa", 3), 3)
	defect.Equal(t, longestSubstring("aabaa", 2), 2)
	defect.Equal(t, longestSubstring("aabaa", 3), 0)
	defect.Equal(t, longestSubstring("aaabb", 3), 3)
	defect.Equal(t, longestSubstring("ababbc", 2), 5)
}
