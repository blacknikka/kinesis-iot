package get

import (
	"gorm.io/gorm"
)

type getSummaryStats struct {
	db *gorm.DB
}

func NewGetSummaryStats(db *gorm.DB) *getSummaryStats {
	return &getSummaryStats{
		db: db,
	}
}

func (stats *getSummaryStats) GetSummaryStartAmount(string) (int64, error) {
	return 0, nil
}
