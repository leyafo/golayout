
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

func NewDefaultOption(debug bool, logPath string)(*Options, error){
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
	if debug == true{
		l := zap.NewAtomicLevel()
		l.SetLevel(zapcore.DebugLevel)
		opt.Level = l
	}

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
	return opt, nil
}

func InitLog(opts *Options)(err error){
	zapLogger, err = opts.Build()
	if err != nil{
		return err
	}
	sugaredLogger = zapLogger.Sugar()
	return nil
}