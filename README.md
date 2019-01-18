# target-filter
target filter

## Build

* [![Build Status](https://travis-ci.org/go-trellis/filters.png)](https://travis-ci.org/go-trellis/filters)

## a filter tool for customer's filter functions 


## Usage


### repo functions

```golang
// FilterRepo filter repo
type FilterRepo interface {
	// Set filter timeout: is second, and must be above zero.
	SetFilterTimeout(int) error
	// add a filter function : CompareFunc
	AddCompareFunc(name string, cf CompareFunc)
	// remove a compare funcation by name
	RemoveCompareFunc(name string)
	// get filter function
	GetCompareFunc(name string) CompareFunc
	// compare input values and filter values
	Compare(filter *FilterParams, input, target FilterValues) (bool, error)
}

type FilterParams struct {
	// filter names
	Names []string
	// compare type: Sequence (0) or Consistent (1)
	Type  CompareType
}
```


### new a filter

```golang
	manager := filter.New()
	manager.AddFilterFunc(filterName, filterNameFunction)
```


### do sequence filter

```golang
fParams := &filter.FilterParams{Names: []string{"filterName"}
manager.Compare(fParams, nil, nil)
```

Or

```golang
fParams := &filter.FilterParams{Type: filter.CompareTypeSequence, Names: []string{"filterName"}
manager.Compare(fParams, nil, nil)
```

### do consistent filter

```golang
fParams := &filter.FilterParams{Type: filter.CompareTypeConsistent, Names: []string{"filterName"}
manager.Compare(fParams, nil, nil)
```

### Get more

[test_file](filter_test.go)