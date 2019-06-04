package errors

// 错误码及对应信息的结构体
type Error struct {
	ErrCode int    // 错误码
	ErrMsg  string // 对应错误信息
}

//  初始化Error结构体函数
func NewError(code int, msg string) *Error {
	return &Error{ErrCode: code, ErrMsg: msg}
}

// Error结构体方法，返回错误信息
func (err *Error) Error() string {
	return err.ErrMsg
}