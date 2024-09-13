package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/common"
	"pandora/pkg/response"
	"runtime/debug"
)

// 通过接收异常，响应用户请求的中间件
func Exception(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			// 使用断言判断错误是否是用户定义的响应异常
			resp, ok := err.(response.Response)
			if !ok {
				// 如果不是用户想要抛出的异常
				common.SystemLog.Error(err)
				common.SystemLog.Error(string(debug.Stack()))

				// 生成异常响应结构体
				resp = response.Response{
					Code:    response.ServerError,
					Message: response.CustomMessage[response.ServerError],
					Data:    map[string]interface{}{},
				}
			}

			// 响应用户请求，不管正确错误 code 都是 200，实际报错通过响应的数据决定
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
