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

	stats := current.NewCurrentStats(mongoDB)
	count, err := stats.GetCurrentStartAmount("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}
