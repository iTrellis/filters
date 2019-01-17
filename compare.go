// MIT License

// Copyright (c) 2015 rutcode-go

package filters

// TODO
// demension compare method: equal; range; gt; lt; egt ; elt
const (
	_CompareTypeEQUAL = iota
	_CompareTypeRANGE
	_CompareTypeGT
	_CompareTypeEGT
	_CompareTypeLT
	_CompareTypeELT
)

// Compare 校验
func (p *Manager) Compare(targetName string, targetValues TargetValues, compareValues CompareValues) (
	filtered bool, err error) {
	targetDemensions := p.MapDemensions[targetName]
	if targetDemensions == nil {
		return true, ErrTargetNameNotExists
	} else if len(targetDemensions) == 0 {
		return
	}

	if compareValues == nil {
		return
	}

	for k, v := range compareValues {
		demension := targetDemensions[k]
		if demension == nil {
			return true, ErrInvalidDemension
		}

		if targetValues[k] == v {
			continue
		}

		return true, nil
	}
	return
}
