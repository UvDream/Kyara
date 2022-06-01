package core

import (
	"fmt"
	RotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"server/global"
	"server/utils"
	"time"
)

// Zap 日志记录主函数
func Zap() (logger *zap.Logger) {
	//	先检查文件是否存在
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok {
		fmt.Printf("%s文件夹不存在\n", global.Config.Zap.Director)
		//不存在就创建文件夹
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}
	cores := make([]zapcore.Core, 0, 7)
	debugLevel := getEncoderCore(zap.DebugLevel)
	infoLevel := getEncoderCore(zap.InfoLevel)
	warnLevel := getEncoderCore(zap.WarnLevel)
	errorLevel := getEncoderCore(zap.ErrorLevel)
	dPanicLevel := getEncoderCore(zap.DPanicLevel)
	panicLevel := getEncoderCore(zap.PanicLevel)
	fatalLevel := getEncoderCore(zap.FatalLevel)
	switch global.Config.Zap.Level {
	case "debug", "DEBUG":
		cores = append(cores, debugLevel, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "info", "INFO":
		cores = append(cores, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "warn", "WARN":
		cores = append(cores, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "error", "ERROR":
		cores = append(cores, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "dpanic", "DPANIC":
		cores = append(cores, dPanicLevel, panicLevel, fatalLevel)
	case "panic", "PANIC":
		cores = append(cores, panicLevel, fatalLevel)
	case "fatal", "FATAL":
		cores = append(cores, panicLevel, fatalLevel)
	default:
		cores = append(cores, debugLevel, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	}
	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

//getEncoderCore 获取日志编码器
func getEncoderCore(level zapcore.Level) (core zapcore.Core) {
	writer, err := FileRotateLogs.GetWriteSyncer(level.String())
	fmt.Println(writer, err)
	if err != nil {
		fmt.Printf("日志分割失败:%v", err.Error())
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}
func getEncoder() zapcore.Encoder {
	if global.Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Config.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch global.Config.Zap.EncodeLevel {
	case "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Config.Zap.Prefix + "- 2006/01/02 - 15:04:05.000"))
}

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWrite, err := RotateLogs.New(
		path.Join(global.Config.Zap.Director, "%Y-%m-%d", level+".log"),
		RotateLogs.ForceNewFile(),
		RotateLogs.WithClock(RotateLogs.Local),
		RotateLogs.WithMaxAge(time.Duration(global.Config.Zap.MaxAge)*24*time.Hour), //日志留存时间
		RotateLogs.WithRotationTime(time.Hour*24),
	)
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWrite)), err
	}
	return zapcore.AddSync(fileWrite), err
}
