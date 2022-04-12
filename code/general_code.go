package code

import "github.com/arabot777/arabot-go/pkg/ecode"

const (
	ErrBadParams          = 100001 // 参数错误接口
	ErrInternal           = 100002 // 服务内部错误
	ErrPermissionDenied   = 100003 // 禁止访问
	ErrNotFound           = 100004 // 未找到资源
	ErrServiceUnavailable = 100005 // 服务不可用，(如：被熔断了)
	ErrBadID              = 100006 // ID错误
	ErrBadPageSize        = 100007 // page size错误
	ErrBadPageNum         = 100008 // page number错误
)

/**
	基于业务逻辑错误 10xxxx
	CloudServer:  11xxxx
	Lebesgue:     12xxxx
	Hermite:      13xxxx
	Bohrium:      14xxxx
	管理后台:      15xxxx
    WebShell:     16xxxx
*/

func init() {
	ecode.Register(ErrBadParams, "请求的参数错误", "请求错误", "")
	ecode.Register(ErrInternal, "服务器内部错误", "服务器错误", "")
	ecode.Register(ErrServiceUnavailable, "服务器暂时不可用，请稍后重试", "服务器错误", "")
	ecode.Register(ErrPermissionDenied, "没有权限访问该资源", "缺少权限", "")
	ecode.Register(ErrNotFound, "未找到请求的资源", "not found", "")
	ecode.Register(ErrBadID, "参数错误: 无效的ID", "请求错误", "")
	ecode.Register(ErrBadPageSize, "分页错误: pgSz错误", "请求错误", "")
	ecode.Register(ErrBadPageNum, "分页错误: pgNum错误", "请求错误", "")
}
