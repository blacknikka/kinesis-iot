package store

import (
	"os"
	"time"

	"github.com/blacknikka/kinesis-iot/usecases/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type storeSummaryStats struct {
	MongoUsecase mongo.MongoUsecase

	dbName string
}

func NewStoreSummaryStats(usecase mongo.MongoUsecase) *storeSummaryStats {
	return &storeSummaryStats{
		MongoUsecase: usecase,
		dbName:       os.Getenv("DATABSE_NAME"),
	}
}

func (stats *storeSummaryStats) StoreSummaryStartAmount(version string) error {
	// get "start" count
	startCount, err := stats.countMongo(
		bson.D{
			{"kind", "start"},
			{"ver", version},
		},
	)

	if err != nil {
		return err
	}

	// get "error" count
	errorCount, err := stats.countMongo(
		bson.D{
			{"kind", "error"},
			{"ver", version},
		},
	)

	if err != nil {
		return err
	}

	// get "event" count
	eventCount, err := stats.countMongo(
		bson.D{
			{"kind", "event"},
			{"ver", version},
		},
	)

	if err != nil {
		return err
	}

	insert := bson.D{
		{"start", startCount},
		{"error", errorCount},
		{"event", eventCount},
		{"timestamp", time.Now().Unix()},
	}

	err = stats.writeToMongo(insert)
	if err != nil {
		return nil
	}

	return nil
}

func (stats *storeSummaryStats) countMongo(opt bson.D) (int64, error) {
	count, err := stats.MongoUsecase.CountAll(
		stats.dbName,
		os.Getenv("NORMAL_COLLECTION"),
		opt,
	)

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (stats *storeSummaryStats) writeToMongo(document bson.D) error {
	err := stats.MongoUsecase.InsertOne(
		stats.dbName,
		os.Getenv("SUMMARY_COLLECTION"),
		document,
	)

	if err != nil {
		return err
	}
	return nil
}
