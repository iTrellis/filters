// MIT License

// Copyright (c) 2015 rutcode-go

package filters

import (
	"github.com/go-trellis/config"
)

// InitFiltersTypeFromFile 初始化的方法名称
const (
	InitFiltersTypeFromFile = "file"
)

// InitManagerFunc 初始化方法
type InitManagerFunc func(options map[string]interface{}) (*Manager, error)

// MapInitialTypes 初始化的方法函数集合
var MapInitialTypes = map[string]InitManagerFunc{
	InitFiltersTypeFromFile: InitManagerFromFile,
}

// FilterRepo 过滤Repo
type FilterRepo interface {
	Compare(string, TargetValues, CompareValues) (bool, error)
}

// InitManagerFromFile 通过文件初始化Manager
func InitManagerFromFile(options map[string]interface{}) (manager *Manager, err error) {

	if options == nil {
		return nil, ErrNeedConfigFile
	}

	filenameI := options["filename"]
	filename, ok := filenameI.(string)
	if !ok {
		return nil, ErrNeedConfigFile
	}

	if manager == nil {
		manager = &Manager{
			MapDemensions: make(map[string]map[string]*Demension),
		}
	}

	tds := &TargetDemensions{}
	if err = config.NewSuffixReader().Read(filename, tds); err != nil {
		return
	}

	for _, v := range tds.TargetDemensions {
		nameDemensions := manager.MapDemensions[v.TargetName]
		if nameDemensions == nil {
			nameDemensions = make(map[string]*Demension, 0)
		}

		for _, d := range v.Demensions {
			nameDemensions[d.TargetKey] = &d
		}
		manager.MapDemensions[v.TargetName] = nameDemensions
	}
	return
}
