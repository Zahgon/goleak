package goleak

import (
	"go.uber.org/goleak/internal/stack"
)

type TestingT interface {
	Error(...interface{})
}

func filterStacks(stacks []stack.Stack, skipID int, opts *opts) []stack.Stack {
	_ = "STUB: not implemented"
	return nil
}

func Find(options ...Option) error { _ = "STUB: not implemented"; return nil }

type testHelper interface {
	Helper()
}

func VerifyNone(t TestingT, options ...Option) { _ = "STUB: not implemented"; return }
