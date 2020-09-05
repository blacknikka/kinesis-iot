package current

import (
	"fmt"

	"github.com/blacknikka/kinesis-iot/usecases/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type CurrentStats struct {
	MongoUsecase mongo.MongoUsecase
}

func (stats *CurrentStats) GetCurrentStartAmount() (int64, error) {
	count, err := stats.MongoUsecase.CountAll("sample-database", "col", bson.D{{"kind", "start"}})
	if err != nil {
		fmt.Println(err.Error())
	}

	return count, nil
}
