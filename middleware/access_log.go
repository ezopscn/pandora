package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/common"
	"time"
)

// 访问日志中间件
func AccessLog(ctx *gin.Context) {
	// 请求时间
	startTime := time.Now()
	// 处理请求
	ctx.Next()
	// 结束时间
	endTime := time.Now()
	// 执行耗时
	execTime := startTime.Sub(endTime)
	// 请求方式
	method := ctx.Request.Method
	// 请求地址
	requestURI := ctx.Request.RequestURI
	// 状态码
	status := ctx.Writer.Status()
	// 来源 IP
	clientIP := ctx.ClientIP()

	// 完整的日志
	logStr := fmt.Sprintf("%s\t%s\t%d\t%s\t%s",
		method,
		requestURI,
		status,
		execTime.String(),
		clientIP,
	)

	// 打印日志，OPTIONS 请求使用 DEBUG
	if method == "OPTIONS" {
		common.AccessLog.Debug(logStr)
	} else {
		common.AccessLog.Info(logStr)
	}
}
