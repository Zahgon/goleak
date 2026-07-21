package goleak

import (
	"time"

	"go.uber.org/goleak/internal/stack"
)

type Option interface {
	apply(*opts)
}

const _defaultRetries = 20

type opts struct {
	filters      []func(stack.Stack) bool
	maxRetries   int
	maxSleep     time.Duration
	cleanup      func(int)
	runOnFailure bool
}

func (o *opts) apply(opts *opts) { _ = "STUB: not implemented"; return }

type optionFunc func(*opts)

func (f optionFunc) apply(opts *opts) { _ = "STUB: not implemented"; return }

func IgnoreTopFunction(f string) Option { _ = "STUB: not implemented"; return *new(Option) }

func IgnoreAnyFunction(f string) Option { _ = "STUB: not implemented"; return *new(Option) }

func IgnoreCreatedBy(f string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Cleanup(cleanupFunc func(exitCode int)) Option { _ = "STUB: not implemented"; return *new(Option) }

func IgnoreCurrent() Option { _ = "STUB: not implemented"; return *new(Option) }

func RunOnFailure() Option { _ = "STUB: not implemented"; return *new(Option) }

func maxSleep(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func addFilter(f func(stack.Stack) bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func buildOpts(options ...Option) *opts { _ = "STUB: not implemented"; return nil }

func (o *opts) filter(s stack.Stack) bool { _ = "STUB: not implemented"; return false }

func (o *opts) retry(i int) bool { _ = "STUB: not implemented"; return false }

func isTestStack(s stack.Stack) bool { _ = "STUB: not implemented"; return false }

func isSyscallStack(s stack.Stack) bool { _ = "STUB: not implemented"; return false }

func isStdLibStack(s stack.Stack) bool { _ = "STUB: not implemented"; return false }

func isTraceStack(s stack.Stack) bool { _ = "STUB: not implemented"; return false }

func isDNSResolverStack(s stack.Stack) bool { _ = "STUB: not implemented"; return false }
