// MIT License

// Copyright (c) 2016 go-trellis

package filters

import (
	"github.com/go-trellis/errors"
)

const (
	namespace = "go-trellis::filters"
)

// errors
var (
	ErrInvalidFilterName        = errors.TN(namespace, 1000, "invalid filter name")
	ErrFilterFunctionEqualNil   = errors.TN(namespace, 1001, "filter function should not be nil")
	ErrNotSupportedFilterType   = errors.TN(namespace, 1002, "filter type not supported")
	ErrNotInputParams           = errors.TN(namespace, 1003, "not input params")
	ErrFailedExecFilterFunction = errors.TN(namespace, 1004, "failed exec filter function: {{.err}}")
	ErrFailedExecTimeout        = errors.TN(namespace, 1005, "exec filter function timeout")
	ErrTimeoutMustAboveZero     = errors.TN(namespace, 1006, "set timeout must above zero")
	ErrUnknownCompareType       = errors.TN(namespace, 1007, "unknown compare type")
)
