package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

// Logger methods interface
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	// log with field
	InfoWithField(msg string, f ...Field)
	ErrorWithField(msg string, f ...Field)
}

var (
	GLogger *logger = &logger{}
	GVA_LOG *zap.Logger
)

// Logger
type logger struct {
	sugarLogger *zap.SugaredLogger
	l           *zap.Logger
	// zapcore.Core
	Key      string
	Service  string
	logLevel int
	// Path    string
}

func configure() zapcore.WriteSyncer {
	// pathFolderLog := "logs" + utils.GetSlashOs()

	// if _, err := os.Stat(pathFolderLog); os.IsNotExist(err) {
	// 	err := os.Mkdir(pathFolderLog, os.ModeDir)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// w := zapcore.AddSync(&lumberjack.Logger{
	// 	Filename:   pathFolderLog + "logger.log",
	// 	MaxSize:    10, // megabytes
	// 	MaxBackups: 300,
	// 	MaxAge:     100, // days
	// })

	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		// zapcore.AddSync(w),
	)
}

// App Logger constructor
func Newlogger(mode, level, format, service string) Logger {
	logLevel, exist := loggerLevelMap[level]
	if !exist {
		logLevel = zapcore.DebugLevel
	}

	// todo
	// ap
	var encoderCfg zapcore.EncoderConfig
	if mode == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		fmt.Println("mode other development")
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.TimeKey = "time"
	encoderCfg.NameKey = "name"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeDuration = zapcore.NanosDurationEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	// encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	// encoderCfg.LineEnding = zapcore.DefaultLineEnding

	var encoder zapcore.Encoder
	if format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}
	if service != "" {
		encoder.AddString("service", service)
		// encoder.AddString("traceId", utils.GenerateCorrelationID())
	}
	core := zapcore.NewCore(encoder, configure(), zap.NewAtomicLevelAt(logLevel))
	loggerZap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	GVA_LOG = loggerZap

	sugarLogger := loggerZap.Sugar()
	tempLog := &logger{
		sugarLogger: sugarLogger,
		l:           loggerZap,
		logLevel:    int(logLevel),
		// Core:        core,
	}
	GLogger = tempLog
	return tempLog
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Logger methods

func (l *logger) Debug(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	args = append([]interface{}{l.Key, " "}, args...)
	l.sugarLogger.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	template = fmt.Sprintf("%s %s", l.Key, template)
	l.sugarLogger.Fatalf(template, args...)
}

// log with field
// use zap.Logger
const MsgFormat = "[%s] %s"

func (l *logger) InfoWithField(msg string, f ...Field) {
	if l.Key != "" {
		msg = fmt.Sprintf(MsgFormat, l.Key, msg)
		l.l.Info(msg, f...)
		return
	}
	msg = fmt.Sprintf("%s %s", l.Key, msg)
	l.l.Info(msg, f...)
}

func (l *logger) ErrorWithField(msg string, f ...Field) {
	if l.Key != "" {
		msg = fmt.Sprintf(MsgFormat, l.Key, msg)
		l.l.Error(msg, f...)
		return
	}
	msg = fmt.Sprintf("%s %s", l.Key, msg)
	l.l.Error(msg, f...)
}

func (l *logger) DebugWithField(msg string, f ...Field) {
	if l.Key != "" {
		msg = fmt.Sprintf(MsgFormat, l.Key, msg)
		l.l.Debug(msg, f...)
		return
	}
	msg = fmt.Sprintf("%s %s", l.Key, msg)
	l.l.Debug(msg, f...)
}

func (l *logger) WarnWithField(msg string, f ...Field) {
	if l.Key != "" {
		msg = fmt.Sprintf(MsgFormat, l.Key, msg)
		l.l.Warn(msg, f...)
		return
	}
	msg = fmt.Sprintf("%s %s", l.Key, msg)
	l.l.Warn(msg, f...)
}

func GetFieldsTrace(traceId string) Field {
	return zap.String("traceId", traceId)
}

func GetFieldsReqId(requestId string) Field {
	return zap.String("requestId", requestId)
}

func GetFieldsKafkaMessageType(messageType string) Field {
	return zap.String("type", messageType)
}

func GetFieldsWorkerID(workerID string) Field {
	return zap.String("workerID", workerID)
}
