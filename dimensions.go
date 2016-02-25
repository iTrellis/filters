package target_manager

import (
	cr "github.com/go-rut/config_reader"
)

var (
	targetDemensions *TargetDemensions
)

type TargetDemensions struct {
	TargetDemensions []Demensions `json:"target_demensions"`
}

func NewTargetDemensions() *TargetDemensions {
	if targetDemensions == nil {
		targetDemensions = new(TargetDemensions)
	}
	return targetDemensions
}

type Demensions struct {
	TargetName string      `json:"target_name"`
	Demensions []Demension `json:"demensions"`
}

type Demension struct {
	TargetKey   string `json:"target_key"`
	Description string `json:"description"`
}

func (*Manager) JsonFileReader(filename string) {

	tds := new(TargetDemensions)

	if err := cr.NewConfigReader().JsonFileReader(filename, tds); err != nil {
		panic(err)
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
}
