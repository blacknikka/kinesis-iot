package current

import (
	"fmt"
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
		"col",
		bson.D{
			{"kind", "start"},
			{"ver", version},
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	return count, nil
}
