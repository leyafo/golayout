package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	sugaredLogger *zap.SugaredLogger
	zapLogger *zap.Logger
)


type Options struct{
	*zap.Config
}

//NewDefaultOption initialize the log options, the logPath's directory must create before use it.
func NewDefaultOption(debug bool, logPath string) *Options {
	var logPaths []string
	if logPath == ""{ //log to stdout
		logPaths = append(logPaths, "stdout")
	}else{
		logPaths = append(logPaths, logPath, "stdout")
	}

	opt := new(Options)
	opt.Config = &zap.Config{
		Encoding: "console",
		EncoderConfig: zap.NewProductionEncoderConfig(),
		OutputPaths: logPaths,
	}
	l := zap.NewAtomicLevel()
	if debug == true{
		l.SetLevel(zapcore.DebugLevel)
	}else{
		l.SetLevel(zapcore.InfoLevel)
	}
	opt.Level = l

	opt.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		NameKey: "",
		EncodeName:     zapcore.FullNameEncoder,
		MessageKey:     "message",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(time.RFC3339))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return opt
}

//InitLog initialize the log module
func InitLog(opts *Options)(err error){
	zapLogger, err = opts.Build()
	if err != nil{
		return err
	}
	sugaredLogger = zapLogger.Sugar()
	return nil
}