/*
Copyright © 2016 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package filters

import "github.com/iTrellis/common/errors"

const (
	namespace = "iTrellis::filters"
)

// errors
var (
	ErrInvalidFilterName        = errors.TN(namespace, 1000, "invalid filter name")
	ErrFilterFunctionEqualNil   = errors.TN(namespace, 1001, "filter function should not be nil")
	ErrNotSupportedFilterType   = errors.TN(namespace, 1002, "filter type not supported")
	ErrNotExistsInputParams     = errors.TN(namespace, 1003, "not exists input params")
	ErrNotExistsTargetParams    = errors.TN(namespace, 1004, "not exists target params")
	ErrFailedExecFilterFunction = errors.TN(namespace, 1005, "failed exec filter function: {{.err}}")
	ErrFailedExecTimeout        = errors.TN(namespace, 1006, "exec filter function timeout")
	ErrTimeoutMustAboveZero     = errors.TN(namespace, 1007, "set timeout must above zero")
	ErrUnknownCompareType       = errors.TN(namespace, 1008, "unknown compare type")
)
