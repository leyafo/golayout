package logger

import (
	"bridgeswap/pkg/path"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	sugaredLogger *zap.SugaredLogger
	zapLogger     *zap.Logger
)

type Options struct {
	*zap.Config
}

//If we didn't call Init, use the default configuration
func init(){
	defaultConfig := zap.NewDevelopmentConfig()
	zapLogger, _ = defaultConfig.Build(zap.AddCallerSkip(1))
	sugaredLogger = zapLogger.Sugar()
}


//NewDefaultOption initialize the log options, the logPath's directory must create before use it.
func NewDefaultOption(debug bool, logPath string) *Options {
	var logPaths []string
	if logPath == "" { //log to stdout
		logPaths = append(logPaths, "stdout")
	} else {
		path.MakeParentDir(logPath)
		logPaths = append(logPaths, logPath, "stdout")
	}

	opt := new(Options)

	opt.Config = &zap.Config{
		Encoding:      "console",
		EncoderConfig: zap.NewProductionEncoderConfig(),
		OutputPaths:   logPaths,
	}
	l := zap.NewAtomicLevel()
	if debug == true {
		l.SetLevel(zapcore.DebugLevel)
	} else {
		l.SetLevel(zapcore.InfoLevel)
	}
	opt.Level = l

	opt.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		CallerKey:   "caller",
		NameKey:     "",
		EncodeName:  zapcore.FullNameEncoder,
		MessageKey:  "message",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(time.RFC3339))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return opt
}

//Init initialize the log module
func Init(opts *Options) (err error) {
	zapLogger, err = opts.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	sugaredLogger = zapLogger.Sugar()
	return nil
}
