package current

type CurrentStatsInterface interface {
	GetCurrentStartAmount(string) (int64, error)
}

type CurrentStatsResponse struct {
	Kind  string `json:"kind"`
	Stats int64  `json:"stats"`
}
