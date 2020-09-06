package get

type GetSummaryStatsUsecase interface {
	GetSummaryStartAmount(string) (int64, error)
}
