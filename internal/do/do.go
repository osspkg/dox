package do

import (
	"fmt"
	"io"
)

type Do struct {
	If     bool
	Result string
}

func If(compare bool) *Do {
	return &Do{
		If: compare,
	}
}

func (v *Do) Then(msg string, args ...interface{}) *Do {
	if v.If {
		v.Result = fmt.Sprintf(msg, args...)
	}
	return v
}

func (v *Do) Else(msg string, args ...interface{}) *Do {
	if !v.If {
		v.Result = fmt.Sprintf(msg, args...)
	}
	return v
}

func (v *Do) Exec(w io.Writer) (err error) {
	_, err = fmt.Fprint(w, v.Result)
	return
}

func Write(msg string, args ...interface{}) *Do {
	return &Do{
		Result: fmt.Sprintf(msg, args...),
	}
}

func Multi(w io.Writer, all ...interface {
	Exec(w io.Writer) (err error)
}) (err error) {
	for _, exec := range all {
		if err = exec.Exec(w); err != nil {
			return
		}
	}
	return
}

type Pack[T interface {
	Exec(w io.Writer) (err error)
}] struct {
	Data []T
}

func (v *Pack[T]) Exec(w io.Writer) (err error) {
	for _, exec := range v.Data {
		if err = exec.Exec(w); err != nil {
			return
		}
	}
	return
}
