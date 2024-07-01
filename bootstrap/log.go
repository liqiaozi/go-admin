package bootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"lixuefei.com/go-admin/common/constants"
	"lixuefei.com/go-admin/common/utils"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/global/logger"
	"os"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

// 初始化日志
func initLog() {
	// 创建日志目录
	createLogDir()
	// 设置日志等级
	setLogLevel()
	// 初始化
	if global.App.Application.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}
	zapCore := getZapCore()
	zapLogger := zap.New(zapCore, options...)
	logger.Log = zapLogger.Sugar()
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(level.String())
	}

	// 设置编码器
	if global.App.Application.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	fileWriteSyncer := getLogWriter()
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)
	return zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriteSyncer, level),
		zapcore.NewCore(encoder, consoleWriteSyncer, level))
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	filename := global.App.Application.Log.Filename
	if filename == "" {
		filename = global.App.Application.AppConfig.AppName + ".log"
	}
	file := &lumberjack.Logger{
		Filename:   global.App.Application.Log.RootDir + constants.File_Separator + filename,
		MaxSize:    global.App.Application.Log.MaxSize,
		MaxBackups: global.App.Application.Log.MaxBackups,
		MaxAge:     global.App.Application.Log.MaxAge,
		Compress:   global.App.Application.Log.Compress,
	}
	return zapcore.AddSync(file)
}

// 设置日志等级
func setLogLevel() {
	switch global.App.Application.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 常见日志目录
func createLogDir() {
	logDir := global.App.Application.Log.RootDir
	if logDir == "" {
		dir, _ := utils.GetCurrentDir()
		logDir = dir + constants.File_Separator + "log"
	}
	if ok, _ := utils.PathExists(logDir); !ok {
		_ = os.Mkdir(logDir, os.ModePerm)
	}
	global.App.Application.Log.RootDir = logDir
}
