package response

// 用户状态码
const (
	OK           = 200
	NotOK        = 400
	Forbidden    = 403
	NotFound     = 404
	ParamError   = 406
	ServerError  = 500
	Unauthorized = 1000
	Unactived    = 1001
)

// 用户状态码对应的错误信息
const (
	OKMessage           = "请求成功"
	NotOKMessage        = "请求失败"
	ForbiddenMessage    = "无权限访问该资源"
	NotFoundMessage     = "请求资源未找到"
	ParamErrorMessage   = "请求参数错误"
	ServerErrorMessage  = "服务器错误，请联系管理员"
	UnauthorizedMessage = "用户未登录"
	UnactivedMessage    = "用户已禁用"
)

// 用户状态码和消息绑定
var CustomMessage = map[int]string{
	OK:           OKMessage,
	NotOK:        NotOKMessage,
	Forbidden:    ForbiddenMessage,
	NotFound:     NotFoundMessage,
	ParamError:   ParamErrorMessage,
	ServerError:  ServerErrorMessage,
	Unauthorized: UnauthorizedMessage,
	Unactived:    UnactivedMessage,
}

// 统一响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 没响应数据
var emptyData = map[string]interface{}{}

// 响应基础方法
func Result(code int, message string, data interface{}) {
	// 通过抛出异常的方式丢给异常中间件
	// 这样做的好处在于即使处理函数因为某些原因触发 panic，也不会让应用退出
	panic(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// 成功
func Success() {
	Result(OK, CustomMessage[OK], emptyData)
}

// 成功，有数据
func SuccessWithData(data interface{}) {
	Result(OK, CustomMessage[OK], data)
}

// 失败
func Failed() {
	Result(NotOK, CustomMessage[NotOK], emptyData)
}

// 失败，有状态码
func FailedWithCode(code int) {
	Result(code, CustomMessage[NotOK], emptyData)
}

// 失败，有失败信息
func FailedWithMessage(message string) {
	Result(NotOK, message, emptyData)
}

// 失败，有状态码和失败信息
func FailedWithCodeAndMessage(code int, message string) {
	Result(code, message, emptyData)
}
