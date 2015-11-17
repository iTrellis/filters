package target_manager

// TODO
// demension compare method: equal; range; gt; lt; egt ; elt
const (
	_COMPARE_TYPE_EQUAL = iota
	_COMPARE_TYPE_RANGE
	_COMPARE_TYPE_GT
	_COMPARE_TYPE_EGT
	_COMPARE_TYPE_LT
	_COMPARE_TYPE_ELT
)

func (p *Manager) Compare(
	targetName string,
	targetValues TargetValues,
	compareValues CompareValues) (
	filtered bool, err error) {
	targetDemensions := p.GetTargetMapDemensions(targetName)
	if targetDemensions == nil {
		return true, ERR_TARGET_NAME_NOT_EXIST
	} else if len(targetDemensions) == 0 {
		return
	}

	if compareValues == nil {
		return
	}

	for k, v := range compareValues {
		demension := targetDemensions[k]
		if demension == nil {
			return true, ERR_INVALID_DEMEMSION
		}

		if targetValues[k] == v {
			continue
		}

		return true, nil
	}
	return
}
