package main

import (
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
			// request exit of goroutine, id == -1 will return on chan block
			prunereq <- id
		}
	}
}
