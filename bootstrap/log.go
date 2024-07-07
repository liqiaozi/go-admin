package bootstrap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"lixuefei.com/go-admin/common/constants"
	"lixuefei.com/go-admin/common/utils/pathutils"
	"lixuefei.com/go-admin/common/utils/stringutils"
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
	fmt.Println("[bootstrap] init log config begin...")
	// 创建日志目录
	createLogDir()
	// 设置日志等级
	setLogLevel()
	// 初始化日志配置
	if global.App.Server.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}
	zapCore := getZapCore()
	zapLogger := zap.New(zapCore, options...)
	logger.Log = zapLogger.Sugar()
	fmt.Println("[bootstrap] init log config end...")
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	// 调整编码器默认配置
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(level.String())
	}

	// 设置编码器
	var encoder zapcore.Encoder
	if global.App.Server.Log.Format == "json" {
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
	filename := global.App.Server.Log.Filename
	if stringutils.IsEmpty(filename) {
		filename = global.App.Server.ServiceInfo.Name + ".log"
	}
	global.App.Server.Log.Filename = filename
	file := &lumberjack.Logger{
		Filename:   global.App.Server.Log.RootDir + constants.File_Separator + filename,
		MaxSize:    global.App.Server.Log.MaxSize,
		MaxBackups: global.App.Server.Log.MaxBackups,
		MaxAge:     global.App.Server.Log.MaxAge,
		Compress:   global.App.Server.Log.Compress,
	}
	return zapcore.AddSync(file)
}

// 设置日志等级
func setLogLevel() {
	switch global.App.Server.Log.Level {
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

// 创建日志目录
func createLogDir() {
	logDir := global.App.Server.Log.RootDir
	if stringutils.IsEmpty(logDir) {
		dir, _ := pathutils.GetCurrentDir()
		logDir = dir + constants.File_Separator + "logs"
	}
	if ok, _ := pathutils.PathExists(logDir); !ok {
		_ = os.Mkdir(logDir, os.ModePerm)
	}
	global.App.Server.Log.RootDir = logDir
}
