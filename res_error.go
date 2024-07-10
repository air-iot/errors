package errors

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var deadLine = fmt.Errorf("执行超时")

func NewResErrorMsg(err error, msg string, a ...any) error {
	if Is(err, context.DeadlineExceeded) {
		return deadLine
	}
	st, ok := status.FromError(err)
	if ok {
		if st.Code() == codes.DeadlineExceeded {
			return deadLine
		}
	}
	if msg == "" {
		return err
	}
	return Wrapf(err, msg, a...)
}

func NewResError(err error) error {
	if Is(err, context.DeadlineExceeded) {
		return deadLine
	}
	st, ok := status.FromError(err)
	if ok {
		if st.Code() == codes.DeadlineExceeded {
			return deadLine
		}
	}
	return err
}
