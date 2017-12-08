package palindrome

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestIsPalindrome(t *testing.T) {
	defect.Equal(t, isPalindrome(11), true)
	defect.Equal(t, isPalindrome(10), false)
	defect.Equal(t, isPalindrome(101), true)
	defect.Equal(t, isPalindrome(12), false)
	defect.Equal(t, isPalindrome(-12), false)
	defect.Equal(t, isPalindrome(1000021), false)
}
