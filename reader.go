package target_manager

type ReaderRepo interface {
	InitFiltersFile(ConfigFile string) error
	InitFiltersDB(ConfigFile string) error
}
