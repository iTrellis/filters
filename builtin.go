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

// defaults function names
const (
	EqualFunctionName = "iTrellis::filters::equals"
)

// CompareEqualsFunc compare values must equal
func CompareEqualsFunc(input, target FilterValues) (filtered bool, err error) {
	if len(target) == 0 {
		return false, nil
	}

	if len(input) == 0 {
		return true, nil
	}

	for k, v := range input {
		if demension := target[k]; demension == nil {
			return true, ErrNotExistsTargetParams.New()
		} else if demension == v {
			continue
		}
		return true, nil
	}

	return false, nil
}
