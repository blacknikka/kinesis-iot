package current

type CurrentStatsUsecase struct {
	Current CurrentStatsInterface
}

func (usecase CurrentStatsUsecase) GetCurrentStats(version string) (*CurrentStatsResponse, error) {
	stats, err := usecase.Current.GetCurrentStartAmount(version)
	if err != nil {
		return nil, err
	}

	return &CurrentStatsResponse{
		Kind:  "current_stats",
		Stats: stats,
	}, nil
}
