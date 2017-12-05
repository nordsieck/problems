package fizz

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestFizzBuzz(t *testing.T) {
	defect.DeepEqual(t, fizzBuzz(15), []string{
		`1`,
		`2`,
		`Fizz`,
		`4`,
		`Buzz`,
		`Fizz`,
		`7`,
		`8`,
		`Fizz`,
		`Buzz`,
		`11`,
		`Fizz`,
		`13`,
		`14`,
		`FizzBuzz`,
	})
}
