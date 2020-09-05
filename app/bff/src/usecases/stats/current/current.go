package current

type CurrentStatsUsecase interface {
	GetCurrentStats(name string) (*StatsResopnse, error)
}

type StatsResopnse struct {
	Name    string      `json:"name"`
	Content interface{} `json:"content"`
}
