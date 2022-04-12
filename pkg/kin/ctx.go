package kin

import (
	"context"
	"github.com/arabot777/arabot-go/code"
	"github.com/arabot777/arabot-go/pkg/ecode"
	"github.com/arabot777/arabot-go/pkg/trace"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const CodeOK = 0

type Context interface {
	Shortcut
	context.Context
}

type Shortcut interface {
	ReplyForbidden()
	ReplyUnauthed()
	Reply(data interface{})
	ReplyErr(code int, hints ...string)
	ReplyRequestErr(hints ...string)
	InternalErr()
	Notfound()
}

type Trace interface {
	StartTrace(kvs ...trace.Entry) (context.Context, trace.FinishFunc)
	StartTraceNamed(name string, kvs ...trace.Entry) (context.Context, trace.FinishFunc)
}

func NewCtx(c *gin.Context) *ctx {
	return &ctx{c}
}

type ctx struct {
	*gin.Context
}

func (*ctx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*ctx) Done() <-chan struct{} {
	return nil
}

func (*ctx) Err() error {
	return nil
}

func (*ctx) Value(key interface{}) interface{} {
	return nil
}

func (c *ctx) ReplyForbidden() {
	c.Status(http.StatusForbidden)
	c.Abort()
}

func (c *ctx) ReplyUnauthed() {
	c.Status(http.StatusUnauthorized)
	c.Abort()
}

func (c *ctx) Reply(data interface{}) {
	c.JSON(http.StatusOK, &Message{
		Code: CodeOK,
		Data: data,
	})
}

func (c *ctx) ReplyOK() {
	c.JSON(http.StatusOK, &Message{
		Code: CodeOK,
	})
}

func (c *ctx) ReplyErr(code int, hints ...string) {
	msg := &MessageError{
		Code:  code,
		Error: ecode.Render(code),
	}
	if len(hints) > 0 {
		msg.Error.Msg = strings.Join(hints, ", ")
	}
	c.JSON(http.StatusOK, msg)
}

func (c *ctx) ReplyRequestErr(hints ...string) {
	msg := &MessageError{
		Code:  code.ErrBadParams,
		Error: ecode.Render(code.ErrBadParams),
	}
	if len(hints) > 0 {
		msg.Error.Msg = strings.Join(hints, ", ")
	}
	c.JSON(http.StatusOK, msg)
}

func (c *ctx) InternalErr() {
	c.JSON(http.StatusOK, &MessageError{
		Code:  code.ErrInternal,
		Error: ecode.Render(code.ErrInternal),
	})
}

func (c *ctx) Notfound() {
	c.JSON(http.StatusOK, &MessageError{
		Code:  code.ErrNotFound,
		Error: ecode.Render(code.ErrNotFound),
	})
}

type Message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type MessageError struct {
	Code  int           `json:"code"`
	Error ecode.ErrInfo `json:"error,omitempty"`
}
