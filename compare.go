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

// FilterValues type map[string]interface{}
type FilterValues map[string]interface{}

// CompareFunc is the function you want to compare input values with target values
// Input Values is the k:v maps with input values
// Target Values is the k:v maps with filter values
// return bool is filtered
// return error if compare function has internal error
// you can write you campare functions like FilterFunc, then manager.AddFilterFunc(name, filterFunc)
type CompareFunc func(input, target FilterValues) (filtered bool, err error)

// CompareType compare type
type CompareType int

// CompareType defines
const (
	CompareTypeSequence CompareType = iota
	CompareTypeConsistent
)

// FilterParams filter params
type FilterParams struct {
	Names []string
	Type  CompareType
}

func (p *FilterParams) valid() (err error) {
	if p == nil {
		return ErrNotExistsInputParams.New()
	}

	if err = p.validType(); err != nil {
		return
	}

	return p.validNames()
}

func (p *FilterParams) validType() error {
	switch p.Type {
	case CompareTypeSequence, CompareTypeConsistent:
		return nil
	}
	return ErrNotSupportedFilterType.New()
}

func (p *FilterParams) validNames() error {
	if len(p.Names) == 0 {
		return ErrInvalidFilterName.New()
	}
	return nil
}
