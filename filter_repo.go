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

import "time"

// FilterRepo filter repo
type FilterRepo interface {
	// Set compare timeout: must be above zero.
	SetCompareTimeout(time.Duration) error
	// add a compare function : CompareFunc
	AddCompareFunc(name string, cf CompareFunc)
	// remove a compare funcation by name
	RemoveCompareFunc(name string)
	// get compare function
	GetCompareFunc(name string) CompareFunc
	// compare input values and filter values
	Compare(filter *FilterParams, input, target FilterValues) (bool, error)
}
