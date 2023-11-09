package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew400Response(t *testing.T) {
	err := New400Response(1, "err1")
	t.Log(err.Error())

	err = Wrap400Err(fmt.Errorf("err2"), 2)
	t.Log(err.Error())

	err = Wrap400Response(fmt.Errorf("err3"), 3, "err4")
	t.Log(err)

	err = Wrap400Err(err, 4)
	t.Log(err.Error())
}

func TestNewError(t *testing.T) {
	e := NewError(1, fmt.Errorf("err1"))
	t.Log(e.Error())
}

func TestNewMsg(t *testing.T) {
	e := NewMsg(1, "a,%d", 2)
	t.Log(e.Error())
}

func TestNewErrorMsg(t *testing.T) {
	e := NewErrorMsg(1, fmt.Errorf("err1"), "a,%d", 2)
	t.Log(e.Error())
	var target error = &ResponseError{}
	t.Log(As(e, &target))
	t.Log(errors.As(e, &target))
}
