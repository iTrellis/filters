// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package errors

import (
	"encoding/json"
)

// ErrorContext map contexts
type ErrorContext map[string]interface{}

func (p ErrorContext) String() string {
	if p == nil {
		return ""
	}

	if bs, e := json.Marshal(p); e == nil {
		return string(bs)
	}
	return ""
}
