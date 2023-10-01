package main

import (
	"github.com/fillu87gyc/lambda-go/lib"
	"github.com/fillu87gyc/lambda-go/lib/zap"
)

func main() {
	zap.InfoSkip("main.go:9", 1)
	zap.GetLogger().Info(lib.ProjectRoot())
}
