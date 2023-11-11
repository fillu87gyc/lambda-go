package zap

import (
	"fmt"
	"github.com/fillu87gyc/lambda-go/lib"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
)

type Logger struct {
	logger *zap.Logger
}

func GetLogger() *zap.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	logger, _ := myConfig(level).Build()
	return logger
}

func JstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "01-02(Mon) 15:04:05.000"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

func myConfig(level zap.AtomicLevel) zap.Config {
	c := zap.Config{
		Level:    level,
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "Msg",
			LevelKey:         "Level",
			TimeKey:          "timestamp",
			NameKey:          "Name",
			CallerKey:        "caller",
			FunctionKey:      "",
			StacktraceKey:    "",
			SkipLineEnding:   false,
			LineEnding:       "",
			EncodeLevel:      zapcore.CapitalColorLevelEncoder,
			EncodeTime:       JstTimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: "",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return c
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
