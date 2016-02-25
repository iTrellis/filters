package target_manager

var (
	targetDemensions *TargetDemensions
)

type TargetDemensions struct {
	TargetDemensions []Demensions `json:"target_demensions"`
}

func (*Manager) NewTargetDemensions() *TargetDemensions {
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
