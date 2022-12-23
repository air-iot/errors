package errors

import (
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
