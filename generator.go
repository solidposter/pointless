package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

func generator(output chan<- datablock, max int, rate int) {
	ts := rand.NewSource(time.Now().UnixNano())
	r := rand.New(ts)
	d := datablock{}

	ticker := time.NewTicker(time.Duration(1/float64(rate)*1000000000) * time.Nanosecond)
	for {
		select {
		case d.timestamp = <-ticker.C:
			d.number = r.Intn(max)
			select {
			case output <- d:
				// Value added to queue
			default:
				// queueu blocked
				// increment queue block counter
				atomic.AddUint64(&randomQblocks, 1)
			}
		}
	}
}
