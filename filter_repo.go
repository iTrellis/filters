// MIT License

// Copyright (c) 2016 go-trellis

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
