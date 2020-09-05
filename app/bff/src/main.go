package main

import (
	"fmt"

	"github.com/blacknikka/kinesis-iot/interfaces/aws/mongo"
	"github.com/blacknikka/kinesis-iot/usecases/stats/current"
)

func main() {
	mongoDB := &mongo.Mongo{}
	if err := mongoDB.Connect(); err != nil {
		fmt.Println(err.Error())
	}

	statas := &current.CurrentStats{
		MongoUsecase: mongoDB,
	}
	count, err := statas.GetCurrentStartAmount()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}
