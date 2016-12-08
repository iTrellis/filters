// MIT License

// Copyright (c) 2015 rutcode-go

package target_manager

import (
	"errors"
)

var (
	ERR_NEED_TARGET_VALUES    = errors.New("need target values")
	ERR_TARGET_NAME_NOT_EXIST = errors.New("target name not exists")
	ERR_INVALID_DEMEMSION     = errors.New("invalid demension")
)
