package errno

import "fmt"

type Errno struct {
	ErrMsg string
}
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e Errno) Error() string {
	return fmt.Sprintf("err_msg:%d", e.ErrMsg)
}

func NewErrNo(msg string) Errno {
	return Errno{
		ErrMsg: msg,
	}
}

func (e Errno) WithMessage(msg string) Errno {
	e.ErrMsg = msg
	return e
}

var (
	PoliceSrvErr = NewErrNo("user service error")
)
