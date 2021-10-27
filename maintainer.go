package main

import (
	"fmt"
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
			select {
			case prunereq <- id:
				// request exit of goroutine
				// dispatcher with ack with a id==-1 on block chan
				fmt.Println("DEBUG: PRUNE BLOCKED FOR ID", id)
			default:
				// prunereq channel blocked
				// try again in one sec, by resetting timer
				timer = time.NewTimer(1 * time.Second)
			}
		}
	}
}
