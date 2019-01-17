// MIT License

// Copyright (c) 2015 rutcode-go

package filters

// TargetDemensions 目标指标
type TargetDemensions struct {
	TargetDemensions []Demensions `json:"target_demensions"`
}

// Demensions 指标组合
type Demensions struct {
	TargetName string      `json:"target_name"`
	Demensions []Demension `json:"demensions"`
}

// Demension 指标
type Demension struct {
	TargetKey   string `json:"target_key"`
	Description string `json:"description"`
}
