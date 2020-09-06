package main

import (
	"fmt"

	"github.com/blacknikka/kinesis-iot/interfaces/aws/mongo"
	"github.com/blacknikka/kinesis-iot/interfaces/stats/summary/get"
	"github.com/blacknikka/kinesis-iot/usecases/stats/current"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	mongoDB := &mongo.Mongo{}
	if err := mongoDB.Connect(); err != nil {
		fmt.Println(err.Error())
	}

	// current stats
	stats := current.NewCurrentStats(mongoDB)
	count, err := stats.GetCurrentStartAmount("ver1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)

	// connect db
	dsn := "host=db user=admin password=admin dbname=db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	getSummary := get.NewGetSummaryStats(db)
	summaryCount, err := getSummary.GetSummaryStartAmount("ver1")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(summaryCount)
}
