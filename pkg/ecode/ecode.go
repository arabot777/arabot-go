package ecode

type ErrInfo struct {
	Title     string
	Msg       string
	Reference string
}

// Register(code, message, title, reference)
func Register(code int, fields ...string) {
	errors.Register(code, fields...)
}

// Render(code)
func Render(code int) ErrInfo {
	return errors.Render(code)
}
