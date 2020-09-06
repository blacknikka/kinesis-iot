package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/blacknikka/kinesis-iot/entities/event"
	"github.com/blacknikka/kinesis-iot/entities/ticker"
	"github.com/blacknikka/kinesis-iot/interfaces/aws/mongo"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/current"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/get"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/store"
	currentUsecase "github.com/blacknikka/kinesis-iot/usecases/stats/current"
	summaryUsecase "github.com/blacknikka/kinesis-iot/usecases/stats/summary/get"
	tickerUsecase "github.com/blacknikka/kinesis-iot/usecases/ticker"
)

func main() {
	mongoDB := &mongo.Mongo{}
	if err := mongoDB.Connect(); err != nil {
		fmt.Println(err.Error())
	}

	// store summary stats
	fmt.Println("store summary")
	storeSummaryStats := store.NewStoreSummaryStats(mongoDB)
	eventFunc := func() error {
		fmt.Println("StoreSummaryStartAmount")
		err := storeSummaryStats.StoreSummaryStartAmount("ver1")
		if err != nil {
			fmt.Printf("error occurred [%v]", err)
		}
		return nil
	}

	// set ticker
	tUsecase := tickerUsecase.TickerUseacse{}
	ctx, cancel := context.WithCancel(context.Background())
	tUsecase.SetTicker(
		ctx,
		&ticker.Ticker{
			Duration: time.Duration(5 * time.Second),
			Event: event.Event{
				Func: eventFunc,
			},
		},
	)

	// get current stats
	fmt.Println("get current stats")
	stats := current.NewCurrentStats(mongoDB)
	currentStatsUsecase := currentUsecase.CurrentStatsUsecase{
		Current: stats,
	}
	count, err := currentStatsUsecase.GetCurrentStats("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)

	// get summary stats
	fmt.Println("get summary stats")
	summaryStats := get.NewGetSummaryStats(mongoDB)
	summaryStatsUsecase := summaryUsecase.SummaryStatsUsecase{
		Summary: summaryStats,
	}
	summaryResult, err := summaryStatsUsecase.GetSummaryStats("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}
	marshaled, err := json.Marshal(summaryResult)
	fmt.Println(string(marshaled))

	time.Sleep(time.Second * 16)
	cancel()
	time.Sleep(time.Second * 2)
}
