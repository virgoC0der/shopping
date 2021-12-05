package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger  *zap.Logger
)

func InitLog() {
	// 日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "/Users/chensx/Desktop/Go/shopping/error.log",
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   false,
	}

	encodingConfig := zap.NewProductionEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(encodingConfig)
	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)),
		zapcore.InfoLevel,
	)

	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	Logger = zap.L()
}
