package main

import (
	"fmt"

	"github.com/blacknikka/kinesis-iot/entities"
)

func main() {
	fmt.Println("hello")

	sample := entities.Record{
		Type: "type",
		Log:  "aaa\nbbb",
	}
	fmt.Println(sample)
}
