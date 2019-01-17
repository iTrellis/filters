// MIT License

// Copyright (c) 2015 rutcode-go

package filters

import (
	"errors"
)

// errors
var (
	ErrNeedConfigFile          = errors.New("file initial type need config path")
	ErrNotFoundInitialFunction = errors.New("initial function not found")
	ErrTargetNameNotExists     = errors.New("target name not exists")
	ErrInvalidDemension        = errors.New("invalid demension")
)
