package ticker

import (
	"context"
	"fmt"
	"time"

	"github.com/blacknikka/kinesis-iot/entities/ticker"
)

type TickerUseacse struct {
	ticker *ticker.Ticker
	base   time.Time
}

func (tu *TickerUseacse) SetTicker(ctx context.Context, t *ticker.Ticker) error {
	tu.ticker = t

	go tu.worker(ctx, tu.ticker)

	return nil
}

func (tu *TickerUseacse) worker(ctx context.Context, t *ticker.Ticker) error {
	tu.base = time.Now()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end ticker")
			return nil
		default:
			elapse := time.Since(tu.base)
			fmt.Println(elapse)
			if elapse.Microseconds() > tu.ticker.Duration.Microseconds() {
				tu.ticker.Event.Func()
				tu.base = time.Now()
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
