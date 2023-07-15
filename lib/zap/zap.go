package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() *zap.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	myConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "Msg",
			LevelKey:         "Level",
			TimeKey:          "",
			NameKey:          "Name",
			CallerKey:        "Caller",
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
