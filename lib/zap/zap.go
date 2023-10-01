package zap

import (
	"fmt"
	"github.com/fillu87gyc/lambda-go/lib"
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

	myConfig := myConfig(level)
	logger, _ := myConfig.Build()
	return logger
}

func myConfig(level zap.AtomicLevel) zap.Config {
	return zap.Config{
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
}

func InfoSkip(msg string, skip int) {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	cfg := myConfig(level)
	cfg.EncoderConfig.EncodeCaller = nil
	logger, _ := cfg.Build()

	_, file, line, _ := runtime.Caller(skip)
	root := lib.ProjectRoot()
	underRoot := file[len(root)+1:]
	m := fmt.Sprintf("%s:%d\t%s", underRoot, line, msg)
	logger.Info(m)
}
