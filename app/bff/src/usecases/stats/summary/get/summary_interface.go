package get

type SummaryStatsInterface interface {
	GetSummaryStartAmount(string) (map[string]interface{}, error)
}

type SummaryStatsResponse struct {
	Kind    string                 `json:"kind"`
	Summary map[string]interface{} `json:"summary"`
}
