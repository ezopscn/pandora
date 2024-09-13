package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"pandora/common"
	"time"
)

// 日志日期格式调整
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(common.MSEC_TIME_FORMAT))
}

// 日志初始化
func NewLogger(cfg common.LoggerConfiguration) *zap.SugaredLogger {
	// 新建配置
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = ZapLocalTimeEncoder          // 调整时间
	config.EncodeLevel = zapcore.CapitalLevelEncoder // 关闭颜色
	var ws zapcore.WriteSyncer                       // 输出
	if cfg.Enabled {
		// 日志文件
		now := time.Now()
		filename := fmt.Sprintf("%s-%04d-%02d-%02d.log",
			cfg.Path,
			now.Year(),
			now.Month(),
			now.Day())

		// 日志切割规则
		hook := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    cfg.MaxSize,
			MaxAge:     cfg.MaxAge,
			MaxBackups: cfg.MaxBackups,
			Compress:   cfg.Compress,
		}

		// 延时关闭
		defer func(hook *lumberjack.Logger) {
			_ = hook.Close()
		}(hook)

		// 输出到控制台和文件
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	} else {
		// 只输出到控制台
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	// 整合日志输出信息
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(config), ws, cfg.Level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}
