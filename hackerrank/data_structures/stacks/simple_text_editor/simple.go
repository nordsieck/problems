package simple

import (
	"fmt"
	"io"
)

const (
	APPEND = iota + 1
	DELETE
	PRINT
	UNDO
)

func ProcessLine(buffer []string, r io.Reader, w io.Writer) []string {
	typ := 0
	fmt.Fscan(r, &typ)

	switch typ {
	case APPEND:
		text := ""
		fmt.Fscan(r, &text)
		return append(buffer, buffer[len(buffer)-1]+text)
	case DELETE:
		i := 0
		fmt.Fscan(r, &i)
		last := buffer[len(buffer)-1]
		return append(buffer, last[:len(last)-i])
	case PRINT:
		i := 0
		fmt.Fscan(r, &i)
		w.Write([]byte{buffer[len(buffer)-1][i-1], '\n'}) // 1 based indexing
		return buffer
	case UNDO:
		return buffer[:len(buffer)-1]
	}
	return nil
}

func Process(r io.Reader, w io.Writer) {
	n := 0
	fmt.Fscan(r, &n)

	buffer := []string{""}
	for i := 0; i < n; i++ {
		buffer = ProcessLine(buffer, r, w)
	}
}
