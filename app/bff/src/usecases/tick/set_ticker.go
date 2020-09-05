package event

import (
	"github.com/blacknikka/kinesis-iot/entities/ticker"
)

type TickerUsecase interface {
	SetTicker(event ticker.Ticker) error
}
