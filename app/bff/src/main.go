package main

import (
	"context"
	"fmt"
	"time"

	"github.com/blacknikka/kinesis-iot/entities/event"
	"github.com/blacknikka/kinesis-iot/entities/ticker"
	"github.com/blacknikka/kinesis-iot/interfaces/aws/mongo"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/current"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/get"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/store"
	"github.com/blacknikka/kinesis-iot/usecases/serve"
	"github.com/blacknikka/kinesis-iot/usecases/serve/handler"
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

	// current stats
	stats := current.NewCurrentStats(mongoDB)
	// summary stats
	summaryStats := get.NewGetSummaryStats(mongoDB)

	// serve
	serveUsecase := serve.ServeUsecase{
		CurrentStatsHandler: handler.ServeCurrentStats{
			CurrentUsecase: stats,
		},
		SummaryStatsHandler: handler.ServeSummaryStats{
			SummaryUsecase: summaryStats,
		},
	}

	if err := serveUsecase.Serve(":8080"); err != nil {
		panic(err.Error())
	}

	// cancel context
	cancel()
	time.Sleep(time.Second * 2)
}
