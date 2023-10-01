package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
)

type Logger struct {
	logger *zap.Logger
}

func GetLogger() *zap.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	myConfig := zap.Config{
		Level:    level,
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "Msg",
			LevelKey:         "Level",
			TimeKey:          "",
			NameKey:          "Name",
			CallerKey:        "caller",
			FunctionKey:      "",
			StacktraceKey:    "",
			SkipLineEnding:   false,
			LineEnding:       "",
			EncodeLevel:      zapcore.CapitalColorLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: "",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ := myConfig.Build()
	return logger
}

func InfoSkip(msg string, skip int) {
	logger := GetLogger()
	_, file, line, _ := runtime.Caller(skip)
	m := fmt.Sprintf("%s:%d %s", file, line, msg)
	logger.Info(m)
}
