package resp

import (
	"fmt"
)

// ErrorResp ： 错误码返回对象
type ErrorResp struct {
	Request   string `json:"request"`
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}

func (errorResp *ErrorResp) String() string {
	return fmt.Sprintf("request=%v\nErrorCode=%v\nError=%v\n", errorResp.Request, errorResp.ErrorCode, errorResp.Error)
}
