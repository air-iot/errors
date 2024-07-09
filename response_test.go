package errors

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
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
	var target = &ResponseError{}
	t.Log(As(e, &target))
	t.Log(errors.As(e, &target))
}

func Test_errorIs(t *testing.T) {
	err1 := fmt.Errorf("测试1")
	err2 := Wrap(err1, "测试2")
	err3 := fmt.Errorf("err2: %w", err1)
	err4 := WithStack(err2)
	t.Log(err2.Error())
	t.Log(err3.Error())
	t.Log(errors.Is(err4, err2))
}

func TestDeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	time.Sleep(time.Second * 2)
	select {
	case <-ctx.Done():
		t.Error(ctx.Err())
		t.Log(errors.Is(ctx.Err(), context.DeadlineExceeded))
	}
}

func Test_errorIs2(t *testing.T) {
	err0 := fmt.Errorf("==========s")
	err1 := Wrap400Err(err0, 123)
	err2 := Wrap(err1, "测试2")
	//err3 := fmt.Errorf("err2: %w", err1)
	err4 := WithStack(err2)
	t.Log(err0.Error())
	t.Log(err1.Error())
	t.Log(errors.Is(err4, err0))
}
