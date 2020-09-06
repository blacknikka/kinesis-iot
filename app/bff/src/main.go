package main

import (
	"encoding/json"
	"fmt"

	"github.com/blacknikka/kinesis-iot/interfaces/aws/mongo"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/current"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/get"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/store"
)

func main() {
	mongoDB := &mongo.Mongo{}
	if err := mongoDB.Connect(); err != nil {
		fmt.Println(err.Error())
	}

	// current stats
	fmt.Println("get current stats")
	stats := current.NewCurrentStats(mongoDB)
	count, err := stats.GetCurrentStartAmount("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)

	// store summary stats
	fmt.Println("store summary")
	storeSummaryStats := store.NewStoreSummaryStats(mongoDB)
	err = storeSummaryStats.StoreSummaryStartAmount("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}

	// summary stats
	fmt.Println("get summary stats")
	summaryStats := get.NewGetSummaryStats(mongoDB)
	summaryResult, err := summaryStats.GetSummaryStartAmount("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}
	marshaled, err := json.Marshal(summaryResult)
	fmt.Println(string(marshaled))
}
