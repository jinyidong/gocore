/**
* @Description: 返回值结构
* @Author: jinyidong
* @Date: 2021/6/17
* @Version V1.0
 */
package util

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

var (
	// OK
	OK           = response(0, "ok")             // 通用成功
	UnKnownError = response(-1, "unknown error") // 通用错误

	// 服务级错误码
	ParseError      = response(-32700, "parse error")
	InvalidRequest  = response(-32600, "invalid request")
	MethodNotFound  = response(-32601, "method not found")
	InvalidParams   = response(-32602, "invalid params")
	InternalError   = response(-32603, "internal error")
	ServerError     = response(-32000, "server error")
	PermissionError = response(-33000, "not permit")
	// ......
)

// 构造函数
func response(code int, msg string) Response {
	return Response{
		Code:    code,
		Message: msg,
		Result:  nil,
	}
}

// 追加响应数据
func (res *Response) WithResult(result interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Result:  result,
	}
}

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Result:  res.Result,
	}
}
