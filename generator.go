package main

//
// Copyright (c) 2021 Tony Sarendal <tony@polarcap.org>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

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
				// Value added to channel
			default:
				// channel blocked
				// increment channel blocked counter
				atomic.AddUint64(&randomQblocks, 1)
			}
		}
	}
}
