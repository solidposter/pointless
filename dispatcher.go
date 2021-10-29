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
	"fmt"
	"time"
)

func dispatcher(blocks <-chan datablock, lifetime int, pruneQsize int) {
	t := make(map[int]chan datablock)
	d := datablock{}
	var numblocks int
	var numprunes int
	var id int

	ticker := time.NewTicker(1 * time.Second)
	prune := make(chan int, pruneQsize)
	for {
		select {
		case d = <-blocks:
			numblocks++
			_, ok := t[d.number]
			if ok == true {
				t[d.number] <- d
			} else {
				m := make(chan datablock, 10)
				go maintainer(d.number, lifetime, m, prune)
				t[d.number] = m
				t[d.number] <- d
			}
		case id = <-prune:
			numprunes++
			t[id] <- datablock{-1, time.Now()}
			delete(t, id)
		case <-ticker.C:
			fmt.Println("Dispatcher values received:", numblocks)
			fmt.Println("Dispatcher prunes received:", numprunes)
			fmt.Println("Dispatcher map size:", len(t))
			numblocks = 0
			numprunes = 0
		}
	}
}
