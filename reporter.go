package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var pruneQblocks uint64
var randomQblocks uint64

func reporter(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("pruneQblocks: ", atomic.LoadUint64(&pruneQblocks))
			fmt.Println("randomQblocks:", atomic.LoadUint64(&randomQblocks))
			// reset the counters to zero
			atomic.StoreUint64(&pruneQblocks, 0)
			atomic.StoreUint64(&randomQblocks, 0)
		}
	}
}
