package current

type CurrentStatsUsecase interface {
	GetCurrentStartAmount() (int64, error)
}
