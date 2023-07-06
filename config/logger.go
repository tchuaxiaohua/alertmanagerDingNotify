package config

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// getEncoder 使用zap库默认字段
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// getLogWriter 定义日志文件
func getLogWriter() zapcore.WriteSyncer {

	lumberWriteSyncer := &lumberjack.Logger{
		Filename:   cfg.Log.FileName,
		MaxSize:    cfg.Log.MaxSize,
		MaxAge:     cfg.Log.MaxAge,
		MaxBackups: cfg.Log.MaxBackups,
		LocalTime:  false,
		Compress:   false,
	}
	return zapcore.AddSync(lumberWriteSyncer)
}

func InitLogger(level string) {
	// 日志级别转换
	l, _ := zapcore.ParseLevel(level)
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writerSyncer, zapcore.AddSync(os.Stdout)), l)

	Log = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Log)
	return
}
