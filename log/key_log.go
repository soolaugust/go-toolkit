package log

import (
	"huashu-iot-connector/consts"
	"huashu-iot-connector/tool"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()
}

func getLogWriter() zapcore.WriteSyncer {
	fileName := consts.LogDirectory + tool.ProgramMD5() + ".log"
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	go func(p string) {
		for {
			fileInfo, _ := os.Stat(p)
			if fileInfo.Size() > consts.MaxLogSize {
				_ = os.Truncate(p, 0)
			}
			time.Sleep(time.Duration(3) * time.Second)
		}
	}(fileName)
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// KDebug 比Debug更快，内存分配次数也更少，但是只支持强类型的结构化日志
func KDebug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// KInfo 比Info更快，内存分配次数也更少，但是只支持强类型的结构化日志
func KInfo(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// KWarn 比Warn更快，内存分配次数也更少，但是只支持强类型的结构化日志
func KWarn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// KError 比Error更快，内存分配次数也更少，但是只支持强类型的结构化日志
func KError(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
