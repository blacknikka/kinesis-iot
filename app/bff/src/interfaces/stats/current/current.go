package current

import (
	"os"

	"github.com/blacknikka/kinesis-iot/usecases/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type currentStats struct {
	MongoUsecase mongo.MongoUsecase

	dbName string
}

func NewCurrentStats(usecase mongo.MongoUsecase) *currentStats {
	return &currentStats{
		MongoUsecase: usecase,
		dbName:       os.Getenv("DATABSE_NAME"),
	}
}

func (stats *currentStats) GetCurrentStartAmount(version string) (int64, error) {
	count, err := stats.MongoUsecase.CountAll(
		stats.dbName,
		os.Getenv("NORMAL_COLLECTION"),
		bson.D{
			{"kind", "start"},
			{"ver", version},
		},
	)
	if err != nil {
		return 0, err
	}

	return count, nil
}
