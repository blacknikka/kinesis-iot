package store

type StoreSummaryStatsUsecase interface {
	StoreSummaryStartAmount(version string) error
}
