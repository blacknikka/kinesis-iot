package mongo

type MongoUsecase interface {
	Connect() error
	CountAll(db string, collection string, where bson.D) (int64, error)
}
