package target_manager

type TragetReaderRepo interface {
	InitFiltersFile(ConfigFile string)
	InitFiltersDB(ConfigFile string)
}
