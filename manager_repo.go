package target_manager

import (
	cr "github.com/go-rut/config_reader"
)

const (
	InitFiltersTypeFromFile = iota
	InitFiltersTypeFromDB
)

type TragetManagerRepo interface {
	InitFilters(filename string, typ int) error

	Compare(string, TargetValues, CompareValues) (bool, error)
}

func (p *Manager) InitFilters(filename string, typ int) (err error) {
	if typ == InitFiltersTypeFromFile {
		err = p.initFiltersFile(filename)
	} else if typ == InitFiltersTypeFromDB {
		err = p.initFiltersDB(filename)
	}

	if err != nil {
		panic(err)
	}

	return
}

func (p *Manager) initFiltersFile(filename string) (err error) {

	tds := p.NewTargetDemensions()

	if err = cr.NewConfigReader().JsonFileReader(filename, tds); err != nil {
		return
	}

	for _, v := range tds.TargetDemensions {
		nameDemensions := mapDemensions[v.TargetName]
		if nameDemensions == nil {
			nameDemensions = make(map[string]*Demension, 0)
		}

		for _, d := range v.Demensions {
			nameDemensions[d.TargetKey] = &d
		}
		mapDemensions[v.TargetName] = nameDemensions
	}

	return
}

// TODO
func (*Manager) initFiltersDB(filename string) (err error) {
	return
}
