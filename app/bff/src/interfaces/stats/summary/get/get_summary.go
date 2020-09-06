package get

import (
	"os"

	"github.com/blacknikka/kinesis-iot/usecases/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type getSummaryStats struct {
	MongoUsecase mongo.MongoUsecase

	dbName string
}

func NewGetSummaryStats(usecase mongo.MongoUsecase) *getSummaryStats {
	return &getSummaryStats{
		MongoUsecase: usecase,
		dbName:       os.Getenv("DATABSE_NAME"),
	}
}

func (stats *getSummaryStats) GetSummaryStartAmount(string) (map[string]interface{}, error) {
	content, err := stats.MongoUsecase.GetLastOne(
		stats.dbName,
		os.Getenv("SUMMARY_COLLECTION"),
		bson.D{},
	)

	if err != nil {
		return nil, err
	}
	return content, nil
}
