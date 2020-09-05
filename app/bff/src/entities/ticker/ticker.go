package ticker

import (
	"time"

	"github.com/blacknikka/kinesis-iot/entities/event"
)

type Ticker struct {
	Duration time.Duration
	Event    event.Event
}
