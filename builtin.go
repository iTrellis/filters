// MIT License

// Copyright (c) 2016 go-trellis

package filters

// defaults function names
const (
	EqualFunctionName = "go-trellis::filters::equals"
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
