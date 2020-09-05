package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoUsecase interface {
	Connect() error
	CountAll(db string, collection string, where bson.D) (int64, error)
}
