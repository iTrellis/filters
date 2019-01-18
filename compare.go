// MIT License

// Copyright (c) 2016 rutcode-go

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
		return ErrNotInputParams.New()
	}

	if err = p.validType(); err != nil {
		return
	}

	return p.validNames()
}

func (p *FilterParams) validType() error {
	if p.Type != CompareTypeSequence &&
		p.Type != CompareTypeConsistent {
		return ErrNotSupportedFilterType.New()
	}
	return nil
}

func (p *FilterParams) validNames() error {
	if len(p.Names) == 0 {
		return ErrInvalidFilterName.New()
	}
	return nil
}
