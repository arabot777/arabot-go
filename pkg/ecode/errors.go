package ecode

var errors Errors

type Errors interface {
	Render(code int) ErrInfo
	Register(code int, fields ...string)
}

type set struct {
	list map[int]ErrInfo
}

func (s *set) Render(code int) ErrInfo {
	return s.list[code]
}

func newErrorSet() Errors {
	s := &set{
		list: make(map[int]ErrInfo),
	}
	return s
}

// Register(300001, "msg", "title", "reference")
// 注册 错误码
func (s *set) Register(code int, fields ...string) {
	fieldsSz := len(fields)
	ei := ErrInfo{}
	switch {
	case fieldsSz > 0:
		ei.Msg = fields[0]
		fallthrough
	case fieldsSz > 1:
		ei.Title = fields[1]
		fallthrough
	case fieldsSz > 2:
		ei.Reference = fields[2]
	}
	s.list[code] = ei
}

func init() {
	errors = newErrorSet()
}
