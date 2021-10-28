package main

import (
	"fmt"
	"time"
)

func dispatcher(blocks <-chan datablock, lifetime int) {
	t := make(map[int]chan datablock)
	d := datablock{}
	var numblocks int = 0
	var id int

	ticker := time.NewTicker(1 * time.Second)
	prune := make(chan int, 10)
	for {
		select {
		case d = <-blocks:
			numblocks++
			_, ok := t[d.number]
			if ok == true {
				t[d.number] <- d
			} else {
				m := make(chan datablock)
				go maintainer(d.number, lifetime, m, prune)
				t[d.number] = m
				t[d.number] <- d
			}
		case id = <-prune:
			// fmt.Println("DEBUG: DELETING", id)
			t[id] <- datablock{-1, time.Now()}
			delete(t, id)
		case <-ticker.C:
			fmt.Println("Dispatcher values received:", numblocks)
			fmt.Println("Dispatcher map size:", len(t))
			numblocks = 0
		}
	}
}
