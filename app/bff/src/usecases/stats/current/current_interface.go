package current

type CurrentStatsUsecase interface {
	GetCurrentStartAmount(string) (int64, error)
}
