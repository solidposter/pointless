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
	"sync/atomic"
	"time"
)

func maintainer(id int, lifetime int, blocks <-chan datablock, prunereq chan<- int) {
	var blocklist []datablock
	d := datablock{}

	timer := time.NewTimer(time.Duration(lifetime) * time.Second)
	for {
		select {
		case d = <-blocks:
			if d.number == -1 {
				// exit goroutine
				return
			} else {
				blocklist = append(blocklist, d)
				timer = time.NewTimer(time.Duration(lifetime) * time.Second)
			}
		case <-timer.C:
			// request exit of gorouting with non-blocking write to pruneq
			// on failure to write, reset the timer to one sec for later try
			select {
			case prunereq <- id:
			default:
				atomic.AddUint64(&pruneQblocks, 1)
				timer = time.NewTimer(1 * time.Second)
			}
		}
	}
}
