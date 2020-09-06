package get

type SummaryStatsUsecase struct {
	Summary SummaryStatsInterface
}

func (usecase SummaryStatsUsecase) GetSummaryStats(version string) (*SummaryStatsResponse, error) {
	summary, err := usecase.Summary.GetSummaryStartAmount(version)
	if err != nil {
		return nil, err
	}

	return &SummaryStatsResponse{
		Kind:    "current_stats",
		Summary: summary,
	}, nil
}
