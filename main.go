package main

import (
	"github.com/fillu87gyc/lambda-go/lib/zap"
)

func main() {
	logger := zap.GetLogger()
	logger.Info("hello world")
}
