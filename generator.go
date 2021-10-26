package main

import (
	"math/rand"
	"time"
)

func generator(output chan<- datablock, max int, rate float64) {
	ts := rand.NewSource(time.Now().UnixNano())
	r := rand.New(ts)
	d := datablock{}

	ticker := time.NewTicker(time.Duration(1/rate*1000000000) * time.Nanosecond)
	for {
		select {
		case d.timestamp = <-ticker.C:
			d.number = r.Intn(max)
			output <- d
		}
	}
}
