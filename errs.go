package errs

import (
	"fmt"
	"io"
)

// New returns an error with code&&err&&msg.
func New(c Code, err error, msg string) error {
	return &errs{c, msg, err, callers()}
}

type errs struct {
	code Code
	msg  string
	error
	*stack
}

func (e *errs) Error() string { return e.msg }

func (e *errs) Code() int { return int(e.code) }

func (e *errs) Unwrap() error { return e.error }

func (e *errs) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.msg)
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.msg)
	case 'q':
		fmt.Fprintf(s, "%q", e.msg)
	}
}
