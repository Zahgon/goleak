package goleak

import (
	"io"
	"os"
)

var (
	_osExit             = os.Exit
	_osStderr io.Writer = os.Stderr
)

type TestingM interface {
	Run() int
}

func VerifyTestMain(m TestingM, options ...Option) { _ = "STUB: not implemented"; return }
