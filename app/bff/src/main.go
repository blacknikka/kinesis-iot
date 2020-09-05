package main

import (
	"fmt"

	"github.com/blacknikka/kinesis-iot/interfaces/aws/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	mongoDB := mongo.Mongo{}
	if err := mongoDB.Connect(); err != nil {
		fmt.Println(err.Error())
	}

	count, err := mongoDB.CountAll("sample-database", "col", bson.D{{"kind", "start"}})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}
