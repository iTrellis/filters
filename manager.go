// MIT License

// Copyright (c) 2015 rutcode-go

package filters

// Manager 管理者对象
type Manager struct {
	MapDemensions map[string]map[string]*Demension
}

// TargetValues 目标参数
type TargetValues map[string]interface{}

// CompareValues 匹配参数
type CompareValues map[string]interface{}

// NewManager 生成管理者对象
func NewManager(initType string, options map[string]interface{}) (FilterRepo, error) {

	initFunc, ok := MapInitialTypes[initType]
	if !ok {
		return nil, ErrNotFoundInitialFunction
	}

	manager, err := initFunc(options)
	if err != nil {
		return nil, err
	}
	return manager, nil
}
