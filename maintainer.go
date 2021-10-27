package main

import (
	//	"fmt"
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
				//		fmt.Println("DEBUG: PRUNE BLOCKED FOR ID", id)
			default:
				timer = time.NewTimer(1 * time.Second)
			}
		}
	}
}
