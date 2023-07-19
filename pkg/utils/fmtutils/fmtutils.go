package fmtutils

import (
	"fmt"
	"io"
)

func MultiFprintf(ws []io.Writer, format string, a ...any) (n int, err error) {
	var merr error
	for _, w := range ws {
		_, err := fmt.Fprintf(w, format, a...)
		if err != nil {
			merr = err
		}
	}
	return 0, merr
}

func MultiPrintf(ws []io.Writer, a ...any) (n int, err error) {
	var merr error
	for _, w := range ws {
		_, err := fmt.Fprint(w, a...)
		if err != nil {
			merr = err
		}
	}
	return 0, merr
}
