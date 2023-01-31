package errors

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ResponseError 定义响应错误
type ResponseError struct {
	Code       int         // 错误码
	Message    string      // 错误消息
	StatusCode int         // 响应状态码
	ERR        error       // 响应错误
	Field      string      // 错误字段
	Data       interface{} // 模版数据
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func NewError(code int, err error, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(msg, args...),
		ERR:     errors.WithStack(err),
	}
	return res
}

// UnWrapResponse 解包响应错误
func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

// WrapResponseWithField 包装响应错误
func WrapResponseWithField(err error, statusCode int, field string, code int, data interface{}, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		ERR:        errors.WithStack(err),
		StatusCode: statusCode,
		Data:       data,
		Field:      field,
	}
	return res
}

// WrapResponse 包装响应错误
func WrapResponse(err error, statusCode, code int, data interface{}, msg string, args ...interface{}) error {
	return WrapResponseWithField(err, statusCode, "", code, data, msg, args...)
}

func Wrap400ResponseWithField(err error, field string, code int, data interface{}, msg string, args ...interface{}) error {
	return WrapResponseWithField(err, http.StatusBadRequest, field, code, data, msg, args...)
}

// Wrap400ErrResponse 包装错误码为400的响应错误
func Wrap400ErrResponse(err error, code int, data interface{}, msg string, args ...interface{}) error {
	return Wrap400ResponseWithField(err, "", code, data, msg, args...)
}

// Wrap400Response 包装错误码为400的响应错误
func Wrap400Response(err error, code int, msg string, args ...interface{}) error {
	return Wrap400ResponseWithField(err, "", code, nil, msg, args...)
}

func Wrap400Err(err error, code int) error {
	return Wrap400ResponseWithField(err, "", code, nil, "")
}

func WrapField(err error, field string) error {
	errTmp, ok := err.(*ResponseError)
	if ok {
		errTmp.Field = field
		return errTmp
	}
	return err
}

func Wrap400ErrWithField(err error, field string, code int) error {
	return Wrap400ResponseWithField(err, field, code, nil, "")
}

// Wrap500Response 包装错误码为500的响应错误
func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, http.StatusInternalServerError, http.StatusInternalServerError, nil, msg, args...)
}

// NewResponse 创建响应错误
func NewResponse(statusCode, code int, data interface{}, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		StatusCode: statusCode,
		Data:       data,
	}
	if res.ERR == nil {
		res.ERR = fmt.Errorf(msg, args...)
	}
	return res
}

// NewResponseWithField 创建响应错误包含错误字段
func NewResponseWithField(statusCode int, field string, code int, data interface{}, msg string, args ...interface{}) error {
	res := &ResponseError{
		StatusCode: statusCode,
		Field:      field,
		Code:       code,
		Data:       data,
		Message:    fmt.Sprintf(msg, args...),
	}
	if res.ERR == nil {
		res.ERR = fmt.Errorf(msg, args...)
	}
	return res
}

func New400ErrResponseWithField(field string, code int, data interface{}, msg string, args ...interface{}) error {
	return NewResponseWithField(http.StatusBadRequest, field, code, data, msg, args...)
}

func New400ResponseWithField(field string, code int, msg string, args ...interface{}) error {
	return New400ErrResponseWithField(field, code, map[string]interface{}{}, msg, args...)
}

// New400ErrResponse 创建错误码为400的响应错误
func New400ErrResponse(code int, data interface{}, msg string, args ...interface{}) error {
	return New400ErrResponseWithField("", code, data, msg, args...)
}

func New400Response(code int, msg string, args ...interface{}) error {
	return New400ErrResponse(code, map[string]interface{}{}, msg, args...)
}

func New500Error(code int, data interface{}) error {
	return NewResponse(http.StatusInternalServerError, code, data, "")
}

// New500Response 创建错误码为500的响应错误
func New500Response(msg string, args ...interface{}) error {
	return NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, map[string]interface{}{}, msg, args...)
}
