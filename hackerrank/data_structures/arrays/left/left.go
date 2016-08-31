package left

import (
	"fmt"
	"io"
)

func readFirst(r io.Reader) (n, d int, err error) {
	_, err = fmt.Fscanf(r, "%d %d\n", &n, &d)
	return n, d, err
}

func readArray(r io.Reader, n int) ([]int, error) {
	s, buff, err := make([]int, n), 0, error(nil)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscanf(r, "%d", &buff)
		if err != nil {
			return nil, err
		}
		s[i] = buff
	}
	return s, nil
}

func writeSlice(w io.Writer, slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		fmt.Fprintf(w, "%v ", slice[i])
	}
	fmt.Fprint(w, slice[len(slice)-1])
}

func writeRotated(w io.Writer, slice []int, rotate int) {
	writeSlice(w, slice[rotate:])
	fmt.Fprint(w, " ")
	writeSlice(w, slice[:rotate])
	fmt.Fprint(w, "\n")
}

func process(r io.Reader, w io.Writer) error {
	n, d, err := readFirst(r)
	if err != nil {
		return err
	}

	slice, err := readArray(r, n)
	if err != nil {
		return err
	}

	writeRotated(w, slice, d)
	return nil
}
