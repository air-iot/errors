package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

// 定义别名
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	Is           = errors.Is
)

// 定义错误
var (
	ErrBadRequest              = New400Response(400, "请求发生错误")
	ErrInvalidParent           = New400Response(400, "无效的父级节点")
	ErrNotAllowDeleteWithChild = New400Response(400, "含有子级，不能删除")
	ErrNotAllowDelete          = New400Response(400, "资源不允许删除")
	ErrInvalidUserName         = New400Response(400, "无效的用户名")
	ErrInvalidPassword         = New400Response(400, "无效的密码")
	ErrInvalidUser             = New400Response(400, "无效的用户")
	ErrUserDisable             = New400Response(400, "用户被禁用，请联系管理员")

	ErrNoPerm          = NewResponse(401, 401, nil, "无访问权限")
	ErrInvalidToken    = NewResponse(401, 401, nil, "令牌失效")
	ErrNotToken        = NewResponse(401, 401, nil, "未找到令牌")
	ErrNotFound        = WrapResponseWithField(fmt.Errorf("资源不存在"), 400, "", 100030002, nil, "资源不存在")
	ErrMethodNotAllow  = NewResponse(405, 405, nil, "方法不被允许")
	ErrTooManyRequests = NewResponse(429, 429, nil, "请求过于频繁")
	ErrInternalServer  = NewResponse(500, 500, nil, "服务器发生错误")
)
